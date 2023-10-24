package domain_todo

import domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"

type todoEntityCollection []domainTodoPort.TodoEntity

func (c *todoEntityCollection) Append(entity domainTodoPort.TodoEntity) {
	*c = append(*c, entity)
}

func (c *todoEntityCollection) ToSlice() []domainTodoPort.TodoEntity {
	slice := make([]domainTodoPort.TodoEntity, len(*c))
	for i, e := range *c {
		slice[i] = e
	}

	return slice
}
