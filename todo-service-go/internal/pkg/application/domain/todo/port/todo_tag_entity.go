package domain_todo_port

type TodoTagEntity interface {
	ID() string
	TodoID() string
	Key() string
}
