package driven_adapter_todo

import domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"

func normalizeTodoEntity(
	source domainTodoPort.TodoEntity,
) *map[string]interface{} {
	var target map[string]interface{}
	setStringEntryIntoMap(&target, "id", source.ID())
	setStringEntryIntoMap(&target, "user_id", source.UserID())
	setStringEntryIntoMap(&target, "label", source.Label())
	todoTagEntityCollection := source.Tags()
	if todoTagEntityCollection == nil {
		return &target
	}
	var todoTagMapCollection []*map[string]interface{}
	for _, todoTagEntity := range *todoTagEntityCollection {
		todoTagMapCollection = append(todoTagMapCollection, normalizeTodoTagEntity(todoTagEntity))
	}
	target["tags"] = todoTagMapCollection

	return &target
}
