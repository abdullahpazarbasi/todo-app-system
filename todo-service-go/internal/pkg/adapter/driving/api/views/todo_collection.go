package driving_adapter_api_views

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type TodoCollection []*Todo

func NewTodoCollectionFromEntityCollection(
	todoEntityCollection *[]domainTodoPort.TodoEntity,
) *TodoCollection {
	todoCollection := TodoCollection{}
	for _, todoEntity := range *todoEntityCollection {
		todoCollection = append(todoCollection, NewTodoFromEntity(todoEntity))
	}

	return &todoCollection
}

func (c *TodoCollection) Size() int {
	return len(*c)
}
