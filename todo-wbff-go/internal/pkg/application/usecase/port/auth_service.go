package usecase_port

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

type AuthService interface {
	ClaimTokenByCredentials(username, password string) (string, domainFaultPort.Fault)
}
