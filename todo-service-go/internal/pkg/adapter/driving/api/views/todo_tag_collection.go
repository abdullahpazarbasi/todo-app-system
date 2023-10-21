package driving_adapter_api_views

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type TodoTagCollection []*TodoTag

func (c *TodoTagCollection) Keys() []string {
	todoTagKeys := make([]string, 0)
	for _, todoTag := range *c {
		todoTagKeys = append(todoTagKeys, todoTag.Key)
	}

	return todoTagKeys
}

func NewTodoTagCollectionFromEntityCollection(
	todoTagEntityCollection *[]domainTodoPort.TodoTagEntity,
) *TodoTagCollection {
	todoTagCollection := TodoTagCollection{}
	for _, todoTagEntity := range *todoTagEntityCollection {
		todoTagCollection = append(todoTagCollection, NewTodoTagFromEntity(todoTagEntity))
	}

	return &todoTagCollection
}
