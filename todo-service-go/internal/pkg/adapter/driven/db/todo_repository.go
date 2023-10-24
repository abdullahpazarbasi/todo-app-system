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
	for _, tag := range todo.Tags().ToSlice() {
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
	domainTodoPort.TodoEntityCollection,
	error,
) {
	var err error
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

	var todoTagEntity domainTodoPort.TodoTagEntity
	var todoEntityCollection domainTodoPort.TodoEntityCollection
	var todoTagEntityCollection domainTodoPort.TodoTagEntityCollection
	todoEntityCollection, err = r.todoFactory.CreateTodoEntityCollection()
	if err != nil {
		return nil, r.f.WrapError(err)
	}
	for _, todoItem := range todoItems {
		todoTagEntityCollection, err = r.todoFactory.CreateTodoTagEntityCollection()
		if err != nil {
			return nil, r.f.WrapError(err)
		}
		for _, todoTagItem := range todoItem.Tags {
			todoTagEntity, err = r.todoFactory.CreateTodoTagEntity(
				todoTagItem.ID,
				nil,
				todoTagItem.Key,
				&todoTagItem.CreatedAt,
				&todoTagItem.UpdatedAt,
			)
			if err != nil {
				return nil, r.f.WrapError(err)
			}
			todoTagEntityCollection.Append(todoTagEntity)
		}
		var userEntity domainUserPort.UserEntity
		userEntity, err = r.userFactory.CreateUserEntity(todoItem.UserID)
		if err != nil {
			return nil, r.f.WrapError(err)
		}
		var todoEntity domainTodoPort.TodoEntity
		todoEntity, err = r.todoFactory.CreateTodoEntity(
			todoItem.ID,
			userEntity,
			todoItem.Label,
			todoTagEntityCollection,
			&todoItem.CreatedAt,
			&todoItem.UpdatedAt,
		)
		if err != nil {
			return nil, r.f.WrapError(err)
		}
		todoEntityCollection.Append(todoEntity)
	}

	return todoEntityCollection, flt
}

func (r *repository) Update(
	ctx context.Context,
	id string,
	manipulate func(currentTodo domainTodoPort.TodoEntity) (newTodo domainTodoPort.TodoEntity, err error),
) error {
	var err error
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
	currentTodoItem := &Todo{
		ID: id,
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).First(&currentTodoItem).Error)
	if flt != nil {
		tx.WithContext(ctx).Rollback()

		return flt
	}
	var currentUserEntity domainUserPort.UserEntity
	currentUserEntity, err = r.userFactory.CreateUserEntity(currentTodoItem.UserID)
	if err != nil {
		tx.WithContext(ctx).Rollback()

		return r.f.WrapError(err)
	}
	var currentTodoTagEntityCollection domainTodoPort.TodoTagEntityCollection
	currentTodoTagEntityCollection, err = r.todoFactory.CreateTodoTagEntityCollection()
	if err != nil {
		tx.WithContext(ctx).Rollback()

		return r.f.WrapError(err)
	}
	var currentTodoTagEntity domainTodoPort.TodoTagEntity
	for _, currentTodoTagItem := range currentTodoItem.Tags {
		currentTodoTagEntity, err = r.todoFactory.CreateTodoTagEntity(
			currentTodoTagItem.ID,
			nil,
			currentTodoTagItem.Key,
			&currentTodoTagItem.CreatedAt,
			&currentTodoTagItem.UpdatedAt,
		)
		if err != nil {
			tx.WithContext(ctx).Rollback()

			return r.f.WrapError(err)
		}
		currentTodoTagEntityCollection.Append(currentTodoTagEntity)
	}
	var currentTodoEntity domainTodoPort.TodoEntity
	currentTodoEntity, err = r.todoFactory.CreateTodoEntity(
		currentTodoItem.ID,
		currentUserEntity,
		currentTodoItem.Label,
		currentTodoTagEntityCollection,
		&currentTodoItem.CreatedAt,
		&currentTodoItem.UpdatedAt,
	)
	if err != nil {
		tx.WithContext(ctx).Rollback()

		return r.f.WrapError(err)
	}

	var newTodoEntity domainTodoPort.TodoEntity
	newTodoEntity, err = manipulate(currentTodoEntity)
	if err != nil {
		tx.WithContext(ctx).Rollback()

		return r.f.WrapError(err)
	}

	newTodoTagItems := make([]TodoTag, 0)
	for _, newTodoTagEntity := range newTodoEntity.Tags().ToSlice() {
		newTodoTagItems = append(newTodoTagItems, TodoTag{
			ID:        newTodoTagEntity.ID(),
			TodoID:    newTodoEntity.ID(),
			Key:       newTodoTagEntity.Key(),
			CreatedAt: *newTodoTagEntity.CreationTime(),
		})
	}
	newTodoItem := &Todo{
		ID:        newTodoEntity.ID(),
		UserID:    newTodoEntity.User().ID(),
		Label:     newTodoEntity.Label(),
		Tags:      newTodoTagItems,
		CreatedAt: *newTodoEntity.CreationTime(),
	}
	flt = r.translateGormErrorToApplicationFault(tx.WithContext(ctx).Save(&newTodoItem).Error)
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
