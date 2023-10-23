package domain_todo

import (
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
)

type todoEntity struct {
	id     string
	userID string
	label  string
	tags   *[]domainTodoPort.TodoTagEntity
}

func (e *todoEntity) ID() string {
	return e.id
}

func (e *todoEntity) UserID() string {
	return e.userID
}

func (e *todoEntity) Label() string {
	return e.label
}

func (e *todoEntity) Tags() *[]domainTodoPort.TodoTagEntity {
	return e.tags
}

func (e *todoEntity) Normalize() map[string]interface{} {
	normalized := make(map[string]interface{})
	normalized["id"] = e.id
	normalized["user_id"] = e.userID
	normalized["label"] = e.label
	todoTagEntityCollection := e.Tags()
	if todoTagEntityCollection == nil {
		return normalized
	}
	var todoTagMapCollection []map[string]interface{}
	for _, se := range *todoTagEntityCollection {
		todoTagMapCollection = append(todoTagMapCollection, se.Normalize())
	}
	normalized["tags"] = todoTagMapCollection

	return normalized
}
