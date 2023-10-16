package domain_todo

import (
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
)

type todoEntity struct {
	id     string
	userID string
	label  string
	tags   *[]domainTodoPort.TodoTagEntity
}

func (e *todoEntity) ID() string {
	return e.id
}

func (e *todoEntity) UserID() string {
	return e.userID
}

func (e *todoEntity) Label() string {
	return e.label
}

func (e *todoEntity) Tags() *[]domainTodoPort.TodoTagEntity {
	return e.tags
}
