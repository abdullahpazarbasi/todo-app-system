package client_interface_rest_api_registrars

import (
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	drivenAppPortsCore "todo-app-wbff/internal/pkg/app/ports/driven/core"
	drivenAppPortsOs "todo-app-wbff/internal/pkg/app/ports/driven/os"
	restApiHandlers "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/handlers"
)

func RegisterTodoAPI(
	e *echo.Echo,
	h restApiHandlers.TodoHandler,
	eva drivenAppPortsOs.EnvironmentVariableAccessor,
) error {
	jwtConfig := echoJWT.Config{
		SigningKey:  []byte(eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTokenSigningKey)),
		TokenLookup: "header:Authorization:Bearer ",
	}
	api := e.Group("/api")
	api.Use(echoJWT.WithConfig(jwtConfig))
	api.GET("/todos", h.GetCollection)
	api.POST("/todos", h.Post)
	api.PATCH("/todos/:id", h.Patch)
	api.DELETE("/todos/:id", h.Delete)

	return nil
}
