package usecase_port

import (
	"context"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
)

type TodoService interface {
	Add(
		ctx context.Context,
		userID string,
		label string,
		tagKeys []string,
	) (
		string,
		domainFaultPort.Fault,
	)
	FindAllForUser(
		ctx context.Context,
		userID string,
	) (
		domainTodoPort.TodoEntityCollection,
		domainFaultPort.Fault,
	)
	Modify(
		ctx context.Context,
		id string,
		userID string,
		label string,
		tagKeys []string,
	) domainFaultPort.Fault
	Remove(
		ctx context.Context,
		id string,
	) domainFaultPort.Fault
}
