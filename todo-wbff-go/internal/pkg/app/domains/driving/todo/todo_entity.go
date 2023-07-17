package driving_app_domains_todo

import (
	appPortsTodoModels "todo-app-wbff/internal/pkg/app/ports/driving/todo"
)

type todoEntity struct {
	id        string
	userID    string
	label     string
	completed bool
}

func NewTodoEntity(id string, userID string, label string, completed bool) appPortsTodoModels.Todo {
	return &todoEntity{
		id:        id,
		userID:    userID,
		label:     label,
		completed: completed,
	}
}

func (t *todoEntity) ID() string {
	return t.id
}

func (t *todoEntity) UserID() string {
	return t.userID
}

func (t *todoEntity) Label() string {
	return t.label
}

func (t *todoEntity) IsCompleted() bool {
	return t.completed
}
