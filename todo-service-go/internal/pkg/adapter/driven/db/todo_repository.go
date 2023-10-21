package driven_adapter_db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	gormClause "gorm.io/gorm/clause"
	gormLogger "gorm.io/gorm/logger"
	"net"
	"net/http"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type repository struct {
	db          *gorm.DB
	userFactory domainUserPort.Factory
	todoFactory domainTodoPort.Factory
	f           domainFaultPort.Factory
}

func NewRepository(
	db *gorm.DB,
	userFactory domainUserPort.Factory,
	todoFactory domainTodoPort.Factory,
	faultFactory domainFaultPort.Factory,
) domainTodoPort.Repository {
	return &repository{
		db:          db,
		userFactory: userFactory,
		todoFactory: todoFactory,
		f:           faultFactory,
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
			TodoID: tag.Todo().ID(),
			Key:    tag.Key(),
		})
	}
	todoItem := Todo{
		ID:     todo.ID(),
		UserID: todo.User().ID(),
		Label:  todo.Label(),
		Tags:   todoTagItems,
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
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Create(&todoItem).Error)
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
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Preload(gormClause.Associations).Where("user_id = ?", userID).Find(&todoItems).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return nil, flt
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Commit().Error)
	if flt != nil {
		return nil, flt
	}

	var todoTagEntities []domainTodoPort.TodoTagEntity
	todoEntities := make([]domainTodoPort.TodoEntity, 0)
	for _, todoMap := range todoItems {
		todoTagEntities = make([]domainTodoPort.TodoTagEntity, 0)
		for _, todoTagMap := range todoMap.Tags {
			todoTagEntities = append(todoTagEntities, r.todoFactory.CreateTodoTagEntity(
				todoTagMap.ID,
				nil,
				todoTagMap.Key,
			))
		}
		todoEntities = append(todoEntities, r.todoFactory.CreateTodoEntity(
			todoMap.ID,
			r.userFactory.CreateUserEntity(todoMap.UserID),
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
	todoItem.UserID = todo.User().ID()
	todoItem.Label = todo.Label()
	for _, tag := range *todo.Tags() {
		for _, todoTagItem := range todoItem.Tags {
			todoTagItem.TodoID = tag.Todo().ID()
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

		return r.f.CreateFault(
			r.f.Cause(err),
			r.f.Type(faultType),
			r.f.ProposedHTTPStatusCode(proposedHTTPStatusCode),
			r.f.Message(err.Message),
		)
	case *net.OpError:
		return r.f.CreateFault(
			r.f.Cause(err),
			r.f.Type(domainFaultPort.FaultTypeConnectionFailure),
			r.f.ProposedHTTPStatusCode(http.StatusServiceUnavailable),
			r.f.Message("connection error"),
		)
	default:
		if errors.Is(err, gormLogger.ErrRecordNotFound) {
			return r.f.WrapError(
				err,
				r.f.Type(domainFaultPort.FaultTypeItemNotFound),
				r.f.ProposedHTTPStatusCode(http.StatusNotFound),
				r.f.Message("item not found"),
			)
		}

		return r.f.WrapError(
			err,
			r.f.Message(fmt.Sprintf("database error %T", err)),
		)
	}
}
