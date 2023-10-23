package domain_todo_port

type TodoTagEntity interface {
	ID() string
	Key() string
	Normalize() map[string]interface{}
}
