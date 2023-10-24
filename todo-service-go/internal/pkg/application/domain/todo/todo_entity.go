package domain_todo

import (
	"time"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type todoEntity struct {
	id               string
	user             domainUserPort.UserEntity
	label            string
	tags             domainTodoPort.TodoTagEntityCollection
	creationTime     *time.Time
	modificationTime *time.Time
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

func (e *todoEntity) Tags() domainTodoPort.TodoTagEntityCollection {
	return e.tags
}

func (e *todoEntity) CreationTime() *time.Time {
	return e.creationTime
}

func (e *todoEntity) ModificationTime() *time.Time {
	return e.modificationTime
}
