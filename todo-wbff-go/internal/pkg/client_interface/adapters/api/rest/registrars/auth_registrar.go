package client_interface_rest_api_registrars

import (
	"github.com/labstack/echo/v4"
	restApiHandlers "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/handlers"
)

func RegisterAuthAPI(e *echo.Echo, tokenClaimHandler restApiHandlers.TokenClaimHandler) error {
	e.POST("/auth/token-claims", tokenClaimHandler.Post)

	return nil
}
