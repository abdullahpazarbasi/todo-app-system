package domain_todo

import (
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type todoEntity struct {
	id    string
	user  domainUserPort.UserEntity
	label string
	tags  *[]domainTodoPort.TodoTagEntity
}

func (e *todoEntity) ID() string {
	return e.id
}

func (e *todoEntity) User() domainUserPort.UserEntity {
	return e.user
}

func (e *todoEntity) Label() string {
	return e.label
}

func (e *todoEntity) Tags() *[]domainTodoPort.TodoTagEntity {
	return e.tags
}
