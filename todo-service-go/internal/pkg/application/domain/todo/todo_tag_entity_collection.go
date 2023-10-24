package domain_todo

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type todoTagEntityCollection []domainTodoPort.TodoTagEntity

func (c *todoTagEntityCollection) Append(entity domainTodoPort.TodoTagEntity) {
	*c = append(*c, entity)
}

func (c *todoTagEntityCollection) FindByKey(key string) domainTodoPort.TodoTagEntity {
	for _, e := range *c {
		if e.Key() == key {
			return e
		}
	}

	return nil
}

func (c *todoTagEntityCollection) ToSlice() []domainTodoPort.TodoTagEntity {
	slice := make([]domainTodoPort.TodoTagEntity, len(*c))
	for i, e := range *c {
		slice[i] = e
	}

	return slice
}
