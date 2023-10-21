package domain_todo_port

type TodoTagEntity interface {
	ID() string
	Todo() TodoEntity
	Key() string
}
