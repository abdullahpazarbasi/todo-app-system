package domain_user

import (
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type factory struct{}

func NewFactory() domainUserPort.Factory {
	return &factory{}
}

func (f *factory) CreateUserEntity(id string) domainUserPort.UserEntity {
	return &userEntity{id: id}
}
