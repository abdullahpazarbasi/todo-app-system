package usecase

import (
	"context"
	"net/http"
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
	usecasePort "todo-app-service/internal/pkg/application/usecase/port"
)

type todoService struct {
	userFactory    domainUserPort.Factory
	todoFactory    domainTodoPort.Factory
	f              domainFaultPort.Factory
	idGenerator    corePort.UUIDGenerator
	todoRepository domainTodoPort.Repository
}

func NewTodoService(
	userFactory domainUserPort.Factory,
	todoFactory domainTodoPort.Factory,
	faultFactory domainFaultPort.Factory,
	idGenerator corePort.UUIDGenerator,
	todoRepository domainTodoPort.Repository,
) usecasePort.TodoService {
	return &todoService{
		userFactory:    userFactory,
		todoFactory:    todoFactory,
		f:              faultFactory,
		idGenerator:    idGenerator,
		todoRepository: todoRepository,
	}
}

func (s *todoService) Add(
	ctx context.Context,
	userID string,
	label string,
	tagKeys []string,
) (
	string,
	domainFaultPort.Fault,
) {
	var err error

	var newUser domainUserPort.UserEntity
	newUser, err = s.userFactory.CreateUserEntity(userID)
	if err != nil {
		return "", s.f.WrapError(
			err,
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("invalid user"),
		)
	}
	todoID := s.idGenerator.GenerateAsString()
	var newTodoTagEntityCollection domainTodoPort.TodoTagEntityCollection
	newTodoTagEntityCollection, err = s.todoFactory.CreateTodoTagEntityCollectionFromKeys(nil, tagKeys)
	if err != nil {
		return "", s.f.WrapError(
			err,
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("invalid todo tag collection"),
		)
	}
	var newTodo domainTodoPort.TodoEntity
	newTodo, err = s.todoFactory.CreateTodoEntity(
		todoID,
		newUser,
		label,
		newTodoTagEntityCollection,
		nil,
		nil,
	)
	if err != nil {
		return "", s.f.WrapError(
			err,
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("invalid todo"),
		)
	}
	err = s.todoRepository.Create(
		ctx,
		newTodo,
	)
	if err != nil {
		return "", s.f.WrapError(err)
	}

	return todoID, nil
}

func (s *todoService) FindAllForUser(
	ctx context.Context,
	userID string,
) (
	domainTodoPort.TodoEntityCollection,
	domainFaultPort.Fault,
) {
	var err error
	var collection domainTodoPort.TodoEntityCollection
	collection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.f.WrapError(err)
	}

	return collection, nil
}

func (s *todoService) Modify(
	ctx context.Context,
	id string,
	userID string,
	label string,
	tagKeysForManipulation []string,
) domainFaultPort.Fault {
	if id == "" {
		return s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("todo ID must be given"),
		)
	}
	err := s.todoRepository.Update(
		ctx,
		id,
		func(currentTodo domainTodoPort.TodoEntity) (newTodo domainTodoPort.TodoEntity, err error) {
			var newUser domainUserPort.UserEntity
			if userID == "" {
				newUser = currentTodo.User()
			} else {
				newUser, err = s.userFactory.CreateUserEntity(userID)
				if err != nil {
					err = s.f.WrapError(
						err,
						s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
						s.f.Message("invalid user"),
					)
					return
				}
			}
			var newLabel string
			if label == "" {
				newLabel = currentTodo.Label()
			} else {
				newLabel = label
			}
			var newTodoTagEntityCollection domainTodoPort.TodoTagEntityCollection
			if len(tagKeysForManipulation) == 0 {
				newTodoTagEntityCollection = currentTodo.Tags()
			} else {
				newTodoTagEntityCollection, _ = s.todoFactory.CreateTodoTagEntityCollection()
				var tagKey string
				var removing bool
				var newTodoTagEntity domainTodoPort.TodoTagEntity
				for _, tagKeyForManipulation := range tagKeysForManipulation {
					tagKey, removing = resolveOriginalKey(tagKeyForManipulation)
					if !removing {
						currentTodoTagEntity := currentTodo.Tags().FindByKey(tagKey)
						if currentTodoTagEntity == nil {
							newTodoTagEntity, err = s.todoFactory.CreateTodoTagEntity(
								s.idGenerator.GenerateAsString(),
								nil,
								tagKey,
								nil,
								nil,
							)
							if err != nil {
								err = s.f.WrapError(
									err,
									s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
									s.f.Message("invalid todo tag"),
								)
								return
							}
							newTodoTagEntityCollection.Append(newTodoTagEntity)
						} else {
							newTodoTagEntityCollection.Append(currentTodoTagEntity)
						}
					}
				}
			}
			newTodo, err = s.todoFactory.CreateTodoEntity(
				id,
				newUser,
				newLabel,
				newTodoTagEntityCollection,
				currentTodo.CreationTime(),
				nil,
			)
			return
		},
	)
	if err != nil {
		return s.f.WrapError(err)
	}

	return nil
}

func (s *todoService) Remove(
	ctx context.Context,
	id string,
) domainFaultPort.Fault {
	err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return s.f.WrapError(err)
	}

	return nil
}
