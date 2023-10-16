package domain_todo_port

type Factory interface {
	CreateTodoEntity(id string, userID string, label string, tags *[]TodoTagEntity) TodoEntity
	CreateTodoTagEntity(id string, key string) TodoTagEntity
	CreateTodoTagEntityCollectionFromKeySlice(keys []string) *[]TodoTagEntity
}
