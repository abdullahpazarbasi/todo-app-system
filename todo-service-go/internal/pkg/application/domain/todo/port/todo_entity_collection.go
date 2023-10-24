package domain_todo_port

type TodoEntityCollection interface {
	Append(entity TodoEntity)
	ToSlice() []TodoEntity
}
