package usecase

import usecasePort "todo-app-wbff/internal/pkg/application/usecase/port"

type todo struct {
	id        string
	userID    string
	label     string
	completed bool
}

func NewTodo(id string, userID string, label string, completed bool) usecasePort.Todo {
	return &todo{
		id:        id,
		userID:    userID,
		label:     label,
		completed: completed,
	}
}

func (t *todo) ID() string {
	return t.id
}

func (t *todo) UserID() string {
	return t.userID
}

func (t *todo) Value() string {
	return t.label
}

func (t *todo) IsCompleted() bool {
	return t.completed
}
