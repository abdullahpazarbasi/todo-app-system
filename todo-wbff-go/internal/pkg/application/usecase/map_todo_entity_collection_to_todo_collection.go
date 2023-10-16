package usecase

import (
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-wbff/internal/pkg/application/usecase/port"
)

func mapTodoEntityCollectionToTodoCollection(source *[]domainTodoPort.TodoEntity) *[]usecasePort.Todo {
	target := make([]usecasePort.Todo, 0)
	var completed bool
	for _, entity := range *source {
		completed = false
		for _, tag := range *entity.Tags() {
			if tag.Key() == "COMPLETED" {
				completed = true
			}
		}
		target = append(target, NewTodo(
			entity.ID(),
			entity.UserID(),
			entity.Label(),
			completed,
		))
	}

	return &target
}
