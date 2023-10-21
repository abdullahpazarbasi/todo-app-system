package domain_todo

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type todoTagEntity struct {
	id   string
	todo domainTodoPort.TodoEntity
	key  string
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
