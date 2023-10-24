package domain_todo

import (
	"fmt"
	"time"
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type factory struct {
	clock       corePort.Clock
	idGenerator corePort.UUIDGenerator
}

func NewFactory(
	clock corePort.Clock,
	idGenerator corePort.UUIDGenerator,
) domainTodoPort.Factory {
	return &factory{
		clock:       clock,
		idGenerator: idGenerator,
	}
}

func (f *factory) CreateTodoEntity(
	id string,
	user domainUserPort.UserEntity,
	label string,
	tags domainTodoPort.TodoTagEntityCollection,
	creationTime *time.Time,
	modificationTime *time.Time,
) (
	domainTodoPort.TodoEntity,
	error,
) {
	if id == "" {
		return nil, fmt.Errorf("id must be given")
	}
	if user == nil {
		return nil, fmt.Errorf("user must be given")
	}
	labelSize := len(label)
	if labelSize == 0 {
		return nil, fmt.Errorf("label must be given")
	}
	if labelSize > 100 {
		return nil, fmt.Errorf("length of label can not be greater than 100")
	}
	if tags == nil {
		tags = &todoTagEntityCollection{}
	}
	if creationTime == nil {
		creationTime = f.clock.Now()
	}
	if modificationTime == nil {
		modificationTime = creationTime
	}
	todo := todoEntity{
		id:               id,
		user:             user,
		label:            label,
		tags:             tags,
		creationTime:     creationTime,
		modificationTime: modificationTime,
	}
	for _, todoTag := range tags.ToSlice() {
		if todoTag.(*todoTagEntity).todo == nil {
			todoTag.(*todoTagEntity).todo = &todo
		}
	}

	return &todo, nil
}

func (f *factory) CreateTodoEntityCollection() (
	domainTodoPort.TodoEntityCollection,
	error,
) {
	return &todoEntityCollection{}, nil
}

func (f *factory) CreateTodoTagEntity(
	id string,
	todo domainTodoPort.TodoEntity,
	key string,
	creationTime *time.Time,
	modificationTime *time.Time,
) (
	domainTodoPort.TodoTagEntity,
	error,
) {
	if id == "" {
		return nil, fmt.Errorf("id must be given")
	}
	keySize := len(key)
	if keySize == 0 {
		return nil, fmt.Errorf("key must be given")
	}
	if keySize > 32 {
		return nil, fmt.Errorf("length of key can not be greater than 32")
	}
	if creationTime == nil {
		creationTime = f.clock.Now()
	}
	if modificationTime == nil {
		modificationTime = creationTime
	}

	return &todoTagEntity{
		id:               id,
		todo:             todo,
		key:              key,
		creationTime:     creationTime,
		modificationTime: modificationTime,
	}, nil
}

func (f *factory) CreateTodoTagEntityCollection() (
	domainTodoPort.TodoTagEntityCollection,
	error,
) {
	return &todoTagEntityCollection{}, nil
}

func (f *factory) CreateTodoTagEntityCollectionFromKeys(
	todo domainTodoPort.TodoEntity,
	keys []string,
) (
	domainTodoPort.TodoTagEntityCollection,
	error,
) {
	var entity domainTodoPort.TodoTagEntity
	var collection domainTodoPort.TodoTagEntityCollection
	var err error

	collection, err = f.CreateTodoTagEntityCollection()
	for _, key := range keys {
		entity, err = f.CreateTodoTagEntity(
			f.idGenerator.GenerateAsString(),
			todo,
			key,
			nil,
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("invalid todo tag entity")
		}
		collection.Append(entity)
	}
	if todo != nil {
		todo.(*todoEntity).tags = collection
	}

	return collection, nil
}
