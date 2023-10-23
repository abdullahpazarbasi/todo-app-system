package domain_todo_port

type Factory interface {
	CreateTodoEntity(id string, userID string, label string, tags *[]TodoTagEntity) TodoEntity
	DenormalizeTodoEntity(normalized map[string]interface{}) TodoEntity
	DenormalizeTodoEntityCollection(normalized *[]map[string]interface{}) *[]TodoEntity
	CreateTodoTagEntity(id string, key string) TodoTagEntity
	DenormalizeTodoTagEntity(normalized map[string]interface{}) TodoTagEntity
	DenormalizeTodoTagEntityCollection(normalized *[]map[string]interface{}) *[]TodoTagEntity
	CreateTodoTagEntityCollectionFromKeySlice(keys []string) *[]TodoTagEntity
}
