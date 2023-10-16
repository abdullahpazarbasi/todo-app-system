package domain_todo

import (
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
)

type factory struct {
}

func NewFactory() domainTodoPort.Factory {
	return &factory{}
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
	key string,
) domainTodoPort.TodoTagEntity {
	return &todoTagEntity{
		id:  id,
		key: key,
	}
}

func (f *factory) CreateTodoTagEntityCollectionFromKeySlice(
	keys []string,
) *[]domainTodoPort.TodoTagEntity {
	collection := make([]domainTodoPort.TodoTagEntity, 0)
	for _, key := range keys {
		collection = append(collection, &todoTagEntity{
			key: key,
		})
	}

	return &collection
}
