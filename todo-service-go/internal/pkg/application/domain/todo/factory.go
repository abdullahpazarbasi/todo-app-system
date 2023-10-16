package domain_todo

import (
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
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
	userID string,
	label string,
	tags *[]domainTodoPort.TodoTagEntity,
) domainTodoPort.TodoEntity {
	return &todoEntity{
		id:     id,
		userID: userID,
		label:  label,
		tags:   tags,
	}
}

func (f *factory) CreateTodoTagEntity(
	id string,
	todoID string,
	key string,
) domainTodoPort.TodoTagEntity {
	return &todoTagEntity{
		id:     id,
		todoID: todoID,
		key:    key,
	}
}

func (f *factory) CreateTodoTagEntityCollectionFromKeys(
	todoID string,
	keys []string,
) *[]domainTodoPort.TodoTagEntity {
	collection := make([]domainTodoPort.TodoTagEntity, 0)
	for _, key := range keys {
		collection = append(collection, &todoTagEntity{
			id:     f.idGenerator.GenerateAsString(),
			todoID: todoID,
			key:    key,
		})
	}

	return &collection
}
