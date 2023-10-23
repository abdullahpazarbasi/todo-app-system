package driving_adapter_api_views

import (
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/application/usecase/port"
)

type Todo struct {
	UserID    string `json:"user_id"`
	ID        string `json:"id"`
	Value     string `json:"value"`
	Completed bool   `json:"completed"`
}

func NewTodoViewFromModel(model drivingAppPortsTodo.Todo) *Todo {
	return &Todo{
		UserID:    model.UserID(),
		ID:        model.ID(),
		Value:     model.Value(),
		Completed: model.IsCompleted(),
	}
}
