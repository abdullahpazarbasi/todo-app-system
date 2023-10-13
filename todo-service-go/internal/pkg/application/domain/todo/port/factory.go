package domain_todo_port

type Factory interface {
	CreateTodoEntity(id string, userID string, label string, tags *[]TodoTagEntity) TodoEntity
	CreateTodoTagEntity(id string, todoID string, key string) TodoTagEntity
	CreateTodoTagEntityCollectionFromKeys(todoID string, keys []string) *[]TodoTagEntity
}
