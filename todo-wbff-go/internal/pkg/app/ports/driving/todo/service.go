package driving_app_ports_todo

import (
	"context"
)

type Service interface {
	Add(ctx context.Context, userID string, value string) (*[]Todo, error)
	FindAll(ctx context.Context, userID string) (*[]Todo, error)
	Modify(ctx context.Context, userID string, id string, value string, completedRaw string) (*[]Todo, error)
	Remove(ctx context.Context, userID string, id string) (*[]Todo, error)
}
