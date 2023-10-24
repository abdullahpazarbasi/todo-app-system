package domain_todo_port

import "time"

type TodoTagEntity interface {
	ID() string
	Todo() TodoEntity
	Key() string
	CreationTime() *time.Time
	ModificationTime() *time.Time
}
