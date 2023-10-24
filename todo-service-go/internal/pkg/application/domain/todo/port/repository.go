package domain_todo_port

import (
	"context"
)

type Repository interface {
	Create(
		ctx context.Context,
		todo TodoEntity,
	) error
	FindAllForUser(
		ctx context.Context,
		userID string,
	) (
		TodoEntityCollection,
		error,
	)
	Update(
		ctx context.Context,
		id string,
		manipulate func(currentTodo TodoEntity) (newTodo TodoEntity, err error),
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
}
