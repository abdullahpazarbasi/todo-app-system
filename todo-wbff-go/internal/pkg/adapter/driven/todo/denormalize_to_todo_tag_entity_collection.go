package driven_adapter_todo

import domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"

func denormalizeToTodoTagEntityCollection(
	target *[]domainTodoPort.TodoTagEntity,
	source *[]map[string]interface{},
	todoFactory domainTodoPort.Factory,
) error {
	for _, todoMap := range *source {
		*target = append(*target, todoFactory.CreateTodoTagEntity(
			extractStringEntryFromMap(todoMap, "id"),
			extractStringEntryFromMap(todoMap, "key"),
		))
	}

	return nil
}
