package domain_todo

import (
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type factory struct {
	idGenerator corePort.UUIDGenerator
}

func NewFactory(idGenerator corePort.UUIDGenerator) domainTodoPort.Factory {
	return &factory{
		idGenerator: idGenerator,
	}
}

func (f *factory) CreateTodoEntity(
	id string,
	user domainUserPort.UserEntity,
	label string,
	tags *[]domainTodoPort.TodoTagEntity,
) domainTodoPort.TodoEntity {
	if id == "" {
		panic("id must be given")
	}
	if user == nil {
		panic("user must be given")
	}
	if label == "" {
		panic("label must be given")
	}
	if tags == nil {
		collection := make([]domainTodoPort.TodoTagEntity, 0)
		tags = &collection
	}
	todo := todoEntity{
		id:    id,
		user:  user,
		label: label,
		tags:  tags,
	}
	for _, todoTag := range *todo.tags {
		if todoTag.(*todoTagEntity).todo == nil {
			todoTag.(*todoTagEntity).todo = &todo
		}
	}

	return &todo
}

func (f *factory) CreateTodoTagEntity(
	id string,
	todo domainTodoPort.TodoEntity,
	key string,
) domainTodoPort.TodoTagEntity {
	if id == "" {
		panic("id must be given")
	}
	if key == "" {
		panic("key must be given")
	}

	return &todoTagEntity{
		id:   id,
		todo: todo,
		key:  key,
	}
}

func (f *factory) CreateTodoTagEntityCollectionFromKeys(
	todo domainTodoPort.TodoEntity,
	keys []string,
) *[]domainTodoPort.TodoTagEntity {
	collection := make([]domainTodoPort.TodoTagEntity, 0)
	for _, key := range keys {
		collection = append(collection, &todoTagEntity{
			id:   f.idGenerator.GenerateAsString(),
			todo: todo,
			key:  key,
		})
	}
	if todo != nil {
		todo.(*todoEntity).tags = &collection
	}

	return &collection
}
