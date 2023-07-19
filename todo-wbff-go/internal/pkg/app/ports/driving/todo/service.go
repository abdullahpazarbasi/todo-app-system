package driving_app_ports_todo

import (
	"context"
	drivingAppPortsError "todo-app-wbff/internal/pkg/app/ports/driving/error"
)

type Service interface {
	Add(
		ctx context.Context,
		userID string,
		value string,
	) (
		*[]Todo,
		drivingAppPortsError.ServiceError,
	)
	FindAll(
		ctx context.Context,
		userID string,
	) (
		*[]Todo,
		drivingAppPortsError.ServiceError,
	)
	Modify(
		ctx context.Context,
		userID string,
		id string,
		value string,
		completedRaw string,
	) (
		*[]Todo,
		drivingAppPortsError.ServiceError,
	)
	Remove(
		ctx context.Context,
		userID string,
		id string,
	) (
		*[]Todo,
		drivingAppPortsError.ServiceError,
	)
}
