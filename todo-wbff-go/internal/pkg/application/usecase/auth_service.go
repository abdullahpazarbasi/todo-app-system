package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"todo-app-wbff/configs"
	corePort "todo-app-wbff/internal/pkg/application/core/port"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
	usecasePort "todo-app-wbff/internal/pkg/application/usecase/port"
)

const tokenLifetimeInMinutes = 10

type authService struct {
	faultFactory                domainFaultPort.Factory
	environmentVariableAccessor corePort.EnvironmentVariableAccessor
}

func NewAuthService(
	faultFactory domainFaultPort.Factory,
	environmentVariableAccessor corePort.EnvironmentVariableAccessor,
) usecasePort.AuthService {
	return &authService{
		faultFactory:                faultFactory,
		environmentVariableAccessor: environmentVariableAccessor,
	}
}

func (s *authService) ClaimTokenByCredentials(username, password string) (string, domainFaultPort.Fault) {
	firstUserUsername := s.environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameFirstUserUsername)
	firstUserPassword := s.environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameFirstUserPassword)
	if username != firstUserUsername || password != firstUserPassword {
		return "", s.faultFactory.CreateFault(
			s.faultFactory.Message("invalid credentials"),
			s.faultFactory.ProposedHTTPStatusCode(http.StatusUnauthorized),
		)
	}

	claims := &jwt.RegisteredClaims{
		Subject:   s.environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameFirstUserID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenLifetimeInMinutes) * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenEncoded, err := token.SignedString([]byte(s.environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameTokenSigningKey)))
	if err != nil {
		return "", s.faultFactory.WrapError(err)
	}

	return tokenEncoded, nil
}
