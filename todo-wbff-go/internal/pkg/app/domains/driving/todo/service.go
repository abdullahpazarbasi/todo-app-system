package driving_app_domains_todo

import (
	"context"
	"strconv"
	drivenAppDomainsTodo "todo-app-wbff/internal/pkg/app/domains/driven/todo"
	drivingAppDomainsError "todo-app-wbff/internal/pkg/app/domains/driving/error"
	drivingAppPortsError "todo-app-wbff/internal/pkg/app/ports/driving/error"
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/app/ports/driving/todo"
)

type service struct {
	todoRepository drivenAppDomainsTodo.Repository
}

func NewService(todoRepository drivenAppDomainsTodo.Repository) *service {
	return &service{
		todoRepository: todoRepository,
	}
}

func (s *service) Add(
	ctx context.Context,
	userID string,
	value string,
) (
	*[]drivingAppPortsTodo.Todo,
	drivingAppPortsError.ServiceError,
) {
	_, err := s.todoRepository.Add(
		ctx,
		NewTodoCandidate(
			userID,
			value,
			false,
		),
	)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "todo could not be added")
	}
	var l *[]*drivenAppDomainsTodo.TodoEntity
	l, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "any todo could not be retrieved")
	}

	return mapTodoEntityCollectionToTodoCollection(l), nil
}

func (s *service) FindAll(
	ctx context.Context,
	userID string,
) (
	*[]drivingAppPortsTodo.Todo,
	drivingAppPortsError.ServiceError,
) {
	var err error
	var l *[]*drivenAppDomainsTodo.TodoEntity
	l, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "any todo could not be retrieved")
	}

	return mapTodoEntityCollectionToTodoCollection(l), nil
}

func (s *service) Modify(
	ctx context.Context,
	userID string,
	id string,
	value string,
	completedRaw string,
) (
	*[]drivingAppPortsTodo.Todo,
	drivingAppPortsError.ServiceError,
) {
	completed, err := strconv.ParseBool(completedRaw)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 400, "malformed parameter 'completed'")
	}

	err = s.todoRepository.Replace(
		ctx,
		NewTodoEntity(
			id,
			userID,
			value,
			completed,
		),
	)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "todo could not be replaced")
	}
	var l *[]*drivenAppDomainsTodo.TodoEntity
	l, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "any todo could not be retrieved")
	}

	return mapTodoEntityCollectionToTodoCollection(l), nil
}

func (s *service) Remove(
	ctx context.Context,
	userID string,
	id string,
) (
	*[]drivingAppPortsTodo.Todo,
	drivingAppPortsError.ServiceError,
) {
	err := s.todoRepository.Remove(
		ctx,
		id,
	)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "todo could not be removed")
	}
	var l *[]*drivenAppDomainsTodo.TodoEntity
	l, err = s.todoRepository.FindAllForUser(ctx, userID)
	if err != nil {
		return nil, drivingAppDomainsError.NewServiceError(err, "", 500, "any todo could not be retrieved")
	}

	return mapTodoEntityCollectionToTodoCollection(l), nil
}

func mapTodoEntityCollectionToTodoCollection(source *[]*drivenAppDomainsTodo.TodoEntity) *[]drivingAppPortsTodo.Todo {
	target := make([]drivingAppPortsTodo.Todo, 0)
	var completed bool
	for _, entity := range *source {
		completed = false
		for _, tag := range entity.Tags {
			if tag == "COMPLETED" {
				completed = true
			}
		}
		target = append(target, NewTodoEntity(
			entity.ID,
			entity.UserID,
			entity.Label,
			completed,
		))
	}

	return &target
}
