package domain_todo_port

type TodoTagEntityCollection interface {
	Append(entity TodoTagEntity)
	FindByKey(key string) TodoTagEntity
	ToSlice() []TodoTagEntity
}
