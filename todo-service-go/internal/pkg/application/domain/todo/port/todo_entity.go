package domain_todo_port

import domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"

type TodoEntity interface {
	ID() string
	User() domainUserPort.UserEntity
	Label() string
	Tags() *[]TodoTagEntity
}
