package domain_user

import (
	"fmt"
	domainUserPort "todo-app-service/internal/pkg/application/domain/user/port"
)

type factory struct{}

func NewFactory() domainUserPort.Factory {
	return &factory{}
}

func (f *factory) CreateUserEntity(id string) (domainUserPort.UserEntity, error) {
	if id == "" {
		return nil, fmt.Errorf("empty user ID")
	}

	return &userEntity{
		id: id,
	}, nil
}
