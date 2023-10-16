package driven_adapter_todo

import domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"

func denormalizeToTodoEntityCollection(
	target *[]domainTodoPort.TodoEntity,
	source *[]map[string]interface{},
	todoFactory domainTodoPort.Factory,
) error {
	var err error

	var todoTagEntityCollection []domainTodoPort.TodoTagEntity
	var todoTagMapCollection []map[string]interface{}
	for _, todoMap := range *source {
		err = denormalizeToTodoTagEntityCollection(
			&todoTagEntityCollection,
			&todoTagMapCollection,
			todoFactory,
		)
		if err != nil {
			return err
		}
		*target = append(*target, todoFactory.CreateTodoEntity(
			extractStringEntryFromMap(todoMap, "id"),
			extractStringEntryFromMap(todoMap, "user_id"),
			extractStringEntryFromMap(todoMap, "label"),
			&todoTagEntityCollection,
		))
	}

	return nil
}
