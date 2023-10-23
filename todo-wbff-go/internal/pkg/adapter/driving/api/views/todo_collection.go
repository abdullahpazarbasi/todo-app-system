package driving_adapter_api_views

import (
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/application/usecase/port"
)

type TodoCollection []*Todo

func (c *TodoCollection) Size() int {
	return len(*c)
}

func NewTodoCollectionFromModelCollection(modelCollection *[]drivingAppPortsTodo.Todo) *TodoCollection {
	collection := TodoCollection{}
	for _, model := range *modelCollection {
		collection = append(collection, NewTodoViewFromModel(model))
	}

	return &collection
}
