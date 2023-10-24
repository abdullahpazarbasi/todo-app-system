package domain_todo_port

import (
	"context"
)

type Repository interface {
	Add(ctx context.Context, todoCandidate TodoEntity) (string, error)
	FindAllForUser(ctx context.Context, userID string) (*[]TodoEntity, error)
	Modify(ctx context.Context, todo TodoEntity) error
	Remove(ctx context.Context, id string) error
}
