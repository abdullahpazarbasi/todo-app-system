package driving_adapter_api_views

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type Todo struct {
	ID     string             `json:"id"`
	UserID string             `json:"user_id"`
	Label  string             `json:"label"`
	Tags   *TodoTagCollection `json:"tags"`
}

func (c *Todo) TagCollection() *TodoTagCollection {
	if c.Tags == nil {
		tags := make(TodoTagCollection, 0)
		c.Tags = &tags
	}

	return c.Tags
}

func NewTodoFromEntity(todoEntity domainTodoPort.TodoEntity) *Todo {
	return &Todo{
		ID:     todoEntity.ID(),
		UserID: todoEntity.User().ID(),
		Label:  todoEntity.Label(),
		Tags:   NewTodoTagCollectionFromEntityCollection(todoEntity.Tags()),
	}
}
