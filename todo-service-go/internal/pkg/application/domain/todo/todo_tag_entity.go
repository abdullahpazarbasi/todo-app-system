package domain_todo

import (
	"time"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
)

type todoTagEntity struct {
	id               string
	todo             domainTodoPort.TodoEntity
	key              string
	creationTime     *time.Time
	modificationTime *time.Time
}

func (e *todoTagEntity) ID() string {
	return e.id
}

func (e *todoTagEntity) Todo() domainTodoPort.TodoEntity {
	return e.todo
}

func (e *todoTagEntity) Key() string {
	return e.key
}

func (e *todoTagEntity) CreationTime() *time.Time {
	return e.creationTime
}

func (e *todoTagEntity) ModificationTime() *time.Time {
	return e.modificationTime
}
