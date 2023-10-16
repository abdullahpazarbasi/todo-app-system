package driven_adapter_db

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
)

type repository struct {
	db           *gorm.DB
	todoFactory  domainTodoPort.Factory
	faultFactory domainFaultPort.Factory
}

func NewRepository(
	db *gorm.DB,
	todoFactory domainTodoPort.Factory,
	faultFactory domainFaultPort.Factory,
) domainTodoPort.Repository {
	return &repository{
		db:           db,
		todoFactory:  todoFactory,
		faultFactory: faultFactory,
	}
}

func (r *repository) Create(
	ctx context.Context,
	todo domainTodoPort.TodoEntity,
) error {
	var flt domainFaultPort.Fault

	todoTagItems := make([]TodoTag, 0)
	for _, tag := range *todo.Tags() {
		todoTagItems = append(todoTagItems, TodoTag{
			ID:     tag.ID(),
			TodoID: tag.TodoID(),
			Key:    tag.Key(),
		})
	}

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if rr := recover(); rr != nil {
			tx.WithContext(ctx).Rollback()
		}
	}()
	flt = r.translateGormErrorToApplicationFault(tx.Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Create(&Todo{
		ID:     todo.ID(),
		UserID: todo.UserID(),
		Label:  todo.Label(),
		Tags:   todoTagItems,
	}).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}

	return r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Commit().Error)
}

func (r *repository) FindAllForUser(
	ctx context.Context,
	userID string,
) (
	*[]domainTodoPort.TodoEntity,
	error,
) {
	var flt domainFaultPort.Fault

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if rr := recover(); rr != nil {
			tx.WithContext(ctx).Rollback()
		}
	}()
	flt = r.translateGormErrorToApplicationFault(tx.Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return nil, flt
	}
	var todoItems []Todo
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Where("user_id = ?", userID).Find(&todoItems).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return nil, flt
	}

	var todoTagEntities []domainTodoPort.TodoTagEntity
	todoEntities := make([]domainTodoPort.TodoEntity, 0)
	for _, todoMap := range todoItems {
		todoTagEntities = make([]domainTodoPort.TodoTagEntity, 0)
		for _, todoTagMap := range todoMap.Tags {
			todoTagEntities = append(todoTagEntities, r.todoFactory.CreateTodoTagEntity(
				todoTagMap.ID,
				todoTagMap.TodoID,
				todoTagMap.Key,
			))
		}
		todoEntities = append(todoEntities, r.todoFactory.CreateTodoEntity(
			todoMap.ID,
			todoMap.UserID,
			todoMap.Label,
			&todoTagEntities,
		))
	}

	return &todoEntities, flt
}

func (r *repository) Update(
	ctx context.Context,
	todo domainTodoPort.TodoEntity,
) error {
	var flt domainFaultPort.Fault

	todoTagItems := make([]TodoTag, 0)
	for _, tag := range *todo.Tags() {
		todoTagItems = append(todoTagItems, TodoTag{
			ID: tag.ID(),
		})
	}
	todoItem := &Todo{
		ID:   todo.ID(),
		Tags: todoTagItems,
	}

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if rr := recover(); rr != nil {
			tx.WithContext(ctx).Rollback()
		}
	}()
	flt = r.translateGormErrorToApplicationFault(tx.Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).First(&todoItem).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}
	todoItem.UserID = todo.UserID()
	todoItem.Label = todo.Label()
	for _, tag := range *todo.Tags() {
		for _, todoTagItem := range todoItem.Tags {
			todoTagItem.TodoID = tag.TodoID()
			todoTagItem.Key = tag.Key()
		}
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Save(&todoItem).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}

	return r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Commit().Error)
}

func (r *repository) Delete(
	ctx context.Context,
	id string,
) error {
	var flt domainFaultPort.Fault

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if rr := recover(); rr != nil {
			tx.WithContext(ctx).Rollback()
		}
	}()
	flt = r.translateGormErrorToApplicationFault(tx.Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Delete(&Todo{ID: id}).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}

	return r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Commit().Error)
}

func (r *repository) translateGormErrorToApplicationFault(gormError error) domainFaultPort.Fault {
	switch err := gormError.(type) {
	case nil:
		return nil
	case *mysql.MySQLError:
		var exists bool
		faultCode := fmt.Sprintf("DBERR%d", err.Number)
		var proposedHTTPStatusCode int
		proposedHTTPStatusCode, exists = mysqlErrorNumberToHTTPStatusCode[err.Number]
		if !exists {
			proposedHTTPStatusCode = http.StatusInternalServerError
		}
		var faultType domainFaultPort.FaultType
		faultType, exists = mysqlErrorNumberToFaultType[err.Number]
		if !exists {
			faultType = domainFaultPort.FaultTypeUnknown
		}

		return r.faultFactory.CreateFault(
			r.faultFactory.Cause(err),
			r.faultFactory.Type(faultType),
			r.faultFactory.Code(faultCode),
			r.faultFactory.ProposedHTTPStatusCode(proposedHTTPStatusCode),
			r.faultFactory.Message(err.Message),
		)
	case *net.OpError:
		return r.faultFactory.CreateFault(
			r.faultFactory.Cause(err),
			r.faultFactory.Type(domainFaultPort.FaultTypeConnectionFailure),
			r.faultFactory.ProposedHTTPStatusCode(http.StatusServiceUnavailable),
			r.faultFactory.Message("connection error"),
		)
	default:
		return r.faultFactory.WrapError(
			err,
			r.faultFactory.Message("database error"),
		)
	}
}
