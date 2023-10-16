package driven_adapter_todo

import domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"

func normalizeTodoTagEntity(source domainTodoPort.TodoTagEntity) *map[string]interface{} {
	var target map[string]interface{}
	setStringEntryIntoMap(&target, "key", source.Key())

	return &target
}
