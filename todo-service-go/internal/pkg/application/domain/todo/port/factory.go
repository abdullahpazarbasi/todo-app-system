package domain_todo_port

import domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"

type Factory interface {
	CreateTodoEntity(
		id string,
		user domainUserPort.UserEntity,
		label string,
		tags *[]TodoTagEntity,
	) TodoEntity
	CreateTodoTagEntity(
		id string,
		todo TodoEntity,
		key string,
	) TodoTagEntity
	CreateTodoTagEntityCollectionFromKeys(
		todo TodoEntity,
		keys []string,
	) *[]TodoTagEntity
}
