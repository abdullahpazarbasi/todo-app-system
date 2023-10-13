package domain_todo_port

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, todo TodoEntity) error
	FindAllForUser(ctx context.Context, userID string) (*[]TodoEntity, error)
	Update(ctx context.Context, todo TodoEntity) error
	Delete(ctx context.Context, id string) error
}
