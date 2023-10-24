package domain_todo_port

import (
	"time"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type Factory interface {
	CreateTodoEntity(
		id string,
		user domainUserPort.UserEntity,
		label string,
		tags TodoTagEntityCollection,
		creationTime *time.Time,
		modificationTime *time.Time,
	) (
		TodoEntity,
		error,
	)
	CreateTodoEntityCollection() (
		TodoEntityCollection,
		error,
	)
	CreateTodoTagEntity(
		id string,
		todo TodoEntity,
		key string,
		creationTime *time.Time,
		modificationTime *time.Time,
	) (
		TodoTagEntity,
		error,
	)
	CreateTodoTagEntityCollection() (
		TodoTagEntityCollection,
		error,
	)
	CreateTodoTagEntityCollectionFromKeys(
		todo TodoEntity,
		keys []string,
	) (
		TodoTagEntityCollection,
		error,
	)
}
