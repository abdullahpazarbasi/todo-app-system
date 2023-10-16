package usecase_port

import (
	"context"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
)

type TodoService interface {
	Add(
		ctx context.Context,
		userID string,
		value string,
	) (
		*[]Todo,
		domainFaultPort.Fault,
	)
	FindAllForUser(
		ctx context.Context,
		userID string,
	) (
		*[]Todo,
		domainFaultPort.Fault,
	)
	Modify(
		ctx context.Context,
		userID string,
		id string,
		value string,
		completedRaw string,
	) (
		*[]Todo,
		domainFaultPort.Fault,
	)
	Remove(
		ctx context.Context,
		userID string,
		id string,
	) (
		*[]Todo,
		domainFaultPort.Fault,
	)
}
