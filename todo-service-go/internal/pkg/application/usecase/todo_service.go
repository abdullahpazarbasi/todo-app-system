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
	if userID == "" {
		return "", s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("user ID must be given"),
		)
	}
	labelSize := len(label)
	if labelSize == 0 {
		return "", s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("label must be given"),
		)
	}
	if labelSize > 100 {
		return "", s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("length of label can not be greater than 100"),
		)
	}
	user := s.userFactory.CreateUserEntity(userID)
	todoID := s.idGenerator.GenerateAsString()
	err := s.todoRepository.Create(
		ctx,
		s.todoFactory.CreateTodoEntity(
			todoID,
			user,
			label,
			s.todoFactory.CreateTodoTagEntityCollectionFromKeys(nil, tagKeys),
		),
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
	*[]domainTodoPort.TodoEntity,
	domainFaultPort.Fault,
) {
	var err error
	var collection *[]domainTodoPort.TodoEntity
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
	tagKeys []string,
) domainFaultPort.Fault {
	if id == "" {
		return s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("todo ID must be given"),
		)
	}
	if userID == "" {
		return s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("user ID must be given"),
		)
	}
	labelSize := len(label)
	if labelSize == 0 {
		return s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("label must be given"),
		)
	}
	if labelSize > 100 {
		return s.f.CreateFault(
			s.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.f.Message("length of label can not be greater than 100"),
		)
	}
	user := s.userFactory.CreateUserEntity(userID)
	err := s.todoRepository.Update(
		ctx,
		s.todoFactory.CreateTodoEntity(
			id,
			user,
			label,
			s.todoFactory.CreateTodoTagEntityCollectionFromKeys(nil, tagKeys),
		),
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
