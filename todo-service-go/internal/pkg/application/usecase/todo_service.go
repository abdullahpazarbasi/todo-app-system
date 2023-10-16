package usecase

import (
	"context"
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-service/internal/pkg/application/usecase/port"
)

type todoService struct {
	todoFactory    domainTodoPort.Factory
	faultFactory   domainFaultPort.Factory
	idGenerator    corePort.UUIDGenerator
	todoRepository domainTodoPort.Repository
}

func NewTodoService(
	todoFactory domainTodoPort.Factory,
	faultFactory domainFaultPort.Factory,
	idGenerator corePort.UUIDGenerator,
	todoRepository domainTodoPort.Repository,
) usecasePort.TodoService {
	return &todoService{
		todoFactory:    todoFactory,
		faultFactory:   faultFactory,
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
	todoID := s.idGenerator.GenerateAsString()
	err := s.todoRepository.Create(
		ctx,
		s.todoFactory.CreateTodoEntity(
			todoID,
			userID,
			label,
			s.todoFactory.CreateTodoTagEntityCollectionFromKeys(todoID, tagKeys),
		),
	)
	if err != nil {
		return "", s.faultFactory.WrapError(err)
	}

	return todoID, nil
}

func (s *todoService) FindAllForUser(
	ctx context.Context,
	userID string,
) (
	*[]domainTodoPort.TodoEntity,
	domainFaultPort.Fault,
) {
	var err error
	var collection *[]domainTodoPort.TodoEntity
	collection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}

	return collection, nil
}

func (s *todoService) Modify(
	ctx context.Context,
	id string,
	userID string,
	label string,
	tagKeys []string,
) domainFaultPort.Fault {
	err := s.todoRepository.Update(
		ctx,
		s.todoFactory.CreateTodoEntity(
			id,
			userID,
			label,
			s.todoFactory.CreateTodoTagEntityCollectionFromKeys(id, tagKeys),
		),
	)
	if err != nil {
		return s.faultFactory.WrapError(err)
	}

	return nil
}

func (s *todoService) Remove(
	ctx context.Context,
	id string,
) domainFaultPort.Fault {
	err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return s.faultFactory.WrapError(err)
	}

	return nil
}
