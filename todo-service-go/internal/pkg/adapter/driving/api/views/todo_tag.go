package driving_adapter_api_views

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type TodoTag struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

func NewTodoTagFromEntity(todoTagEntity domainTodoPort.TodoTagEntity) *TodoTag {
	return &TodoTag{
		ID:  todoTagEntity.ID(),
		Key: todoTagEntity.Key(),
	}
}
