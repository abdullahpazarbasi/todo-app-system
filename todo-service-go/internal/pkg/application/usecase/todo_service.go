package usecase

import (
	"context"
	"todo-app-service/internal/pkg/application/core"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-service/internal/pkg/application/usecase/port"
)

type todoService struct {
	todoFactory    domainTodoPort.Factory
	faultFactory   domainFaultPort.Factory
	todoRepository domainTodoPort.Repository
}

func NewTodoService(
	todoFactory domainTodoPort.Factory,
	faultFactory domainFaultPort.Factory,
	todoRepository domainTodoPort.Repository,
) usecasePort.TodoService {
	return &todoService{
		todoFactory:    todoFactory,
		faultFactory:   faultFactory,
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
	idGenerator := core.ExtractUUIDGeneratorFromContext(ctx)
	todoID := idGenerator.GenerateAsString()
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
		return "", s.faultFactory.WrapError(err, "E891284193050778")
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
		return nil, s.faultFactory.WrapError(err, "E812740732935602")
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
		return s.faultFactory.WrapError(err, "E712643828035880")
	}

	return nil
}

func (s *todoService) Remove(
	ctx context.Context,
	id string,
) domainFaultPort.Fault {
	err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return s.faultFactory.WrapError(err, "E737283570475033")
	}

	return nil
}
