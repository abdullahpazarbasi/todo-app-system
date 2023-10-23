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

func (f *factory) DenormalizeTodoEntity(normalized map[string]interface{}) domainTodoPort.TodoEntity {
	if normalized == nil {
		return nil
	}

	return &todoEntity{
		id:     f.extractStringEntryFromMap(normalized, "id"),
		userID: f.extractStringEntryFromMap(normalized, "user_id"),
		label:  f.extractStringEntryFromMap(normalized, "label"),
		tags:   f.DenormalizeTodoTagEntityCollection(f.extractMapCollectionFromMap(normalized, "tags")),
	}
}

func (f *factory) DenormalizeTodoEntityCollection(normalized *[]map[string]interface{}) *[]domainTodoPort.TodoEntity {
	if normalized == nil {
		return nil
	}
	entityCollection := make([]domainTodoPort.TodoEntity, 0)
	if len(*normalized) == 0 {
		return &entityCollection
	}
	for _, entityMap := range *normalized {
		entityCollection = append(entityCollection, f.DenormalizeTodoEntity(entityMap))
	}

	return &entityCollection
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

func (f *factory) DenormalizeTodoTagEntity(normalized map[string]interface{}) domainTodoPort.TodoTagEntity {
	if normalized == nil {
		return nil
	}

	return &todoTagEntity{
		id:  f.extractStringEntryFromMap(normalized, "id"),
		key: f.extractStringEntryFromMap(normalized, "key"),
	}
}

func (f *factory) DenormalizeTodoTagEntityCollection(normalized *[]map[string]interface{}) *[]domainTodoPort.TodoTagEntity {
	if normalized == nil {
		return nil
	}
	entityCollection := make([]domainTodoPort.TodoTagEntity, 0)
	if len(*normalized) == 0 {
		return &entityCollection
	}
	for _, entityMap := range *normalized {
		entityCollection = append(entityCollection, f.DenormalizeTodoTagEntity(entityMap))
	}

	return &entityCollection
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

func (f *factory) extractStringEntryFromMap(source map[string]interface{}, key string) string {
	rawValue, existent := source[key]
	if !existent {
		return ""
	}
	value, ok := rawValue.(string)
	if !ok {
		return ""
	}

	return value
}

func (f *factory) extractMapCollectionFromMap(source map[string]interface{}, key string) *[]map[string]interface{} {
	rawCollection, existent := source[key]
	if !existent {
		return nil
	}
	emptyMapCollection := make([]map[string]interface{}, 0)
	var collection []interface{}
	var fit bool
	collection, fit = rawCollection.([]interface{})
	if !fit {
		return &emptyMapCollection
	}
	mapCollection := make([]map[string]interface{}, 0)
	var item map[string]interface{}
	for _, rawItem := range collection {
		item, fit = rawItem.(map[string]interface{})
		if !fit {
			return &emptyMapCollection
		}
		mapCollection = append(mapCollection, item)
	}

	return &mapCollection
}
