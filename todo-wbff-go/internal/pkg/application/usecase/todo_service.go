package usecase

import (
	"context"
	"net/http"
	"strconv"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-wbff/internal/pkg/application/usecase/port"
)

type todoService struct {
	faultFactory   domainFaultPort.Factory
	todoFactory    domainTodoPort.Factory
	todoRepository domainTodoPort.Repository
}

func NewTodoService(
	faultFactory domainFaultPort.Factory,
	todoFactory domainTodoPort.Factory,
	todoRepository domainTodoPort.Repository,
) usecasePort.TodoService {
	return &todoService{
		faultFactory:   faultFactory,
		todoFactory:    todoFactory,
		todoRepository: todoRepository,
	}
}

func (s *todoService) Add(
	ctx context.Context,
	userID string,
	value string,
) (
	*[]usecasePort.Todo,
	domainFaultPort.Fault,
) {
	_, err := s.todoRepository.Add(
		ctx,
		s.todoFactory.CreateTodoEntity(
			"",
			userID,
			value,
			s.todoFactory.CreateTodoTagEntityCollectionFromKeySlice([]string{}),
		),
	)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}
	var todoEntityCollection *[]domainTodoPort.TodoEntity
	todoEntityCollection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}

	return mapTodoEntityCollectionToTodoCollection(todoEntityCollection), nil
}

func (s *todoService) FindAllForUser(
	ctx context.Context,
	userID string,
) (
	*[]usecasePort.Todo,
	domainFaultPort.Fault,
) {
	var err error
	var todoEntityCollection *[]domainTodoPort.TodoEntity
	todoEntityCollection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}

	return mapTodoEntityCollectionToTodoCollection(todoEntityCollection), nil
}

func (s *todoService) Modify(
	ctx context.Context,
	userID string,
	id string,
	value string,
	completedRaw string,
) (
	*[]usecasePort.Todo,
	domainFaultPort.Fault,
) {
	completed, err := strconv.ParseBool(completedRaw)
	if err != nil {
		return nil, s.faultFactory.CreateFault(
			s.faultFactory.Cause(err),
			s.faultFactory.ProposedHTTPStatusCode(http.StatusBadRequest),
			s.faultFactory.Message("malformed parameter 'completed'"),
		)
	}
	var todoTagEntityCollection *[]domainTodoPort.TodoTagEntity
	if completed {
		todoTagEntityCollection = s.todoFactory.CreateTodoTagEntityCollectionFromKeySlice([]string{"COMPLETED"})
	}

	err = s.todoRepository.Replace(
		ctx,
		s.todoFactory.CreateTodoEntity(
			id,
			userID,
			value,
			todoTagEntityCollection,
		),
	)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}
	var todoEntityCollection *[]domainTodoPort.TodoEntity
	todoEntityCollection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}

	return mapTodoEntityCollectionToTodoCollection(todoEntityCollection), nil
}

func (s *todoService) Remove(
	ctx context.Context,
	userID string,
	id string,
) (
	*[]usecasePort.Todo,
	domainFaultPort.Fault,
) {
	err := s.todoRepository.Remove(
		ctx,
		id,
	)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}
	var todoEntityCollection *[]domainTodoPort.TodoEntity
	todoEntityCollection, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, s.faultFactory.WrapError(err)
	}

	return mapTodoEntityCollectionToTodoCollection(todoEntityCollection), nil
}
