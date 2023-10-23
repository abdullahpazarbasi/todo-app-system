package domain_todo_port

type TodoEntity interface {
	ID() string
	UserID() string
	Label() string
	Tags() *[]TodoTagEntity
	Normalize() map[string]interface{}
}
