package driving_adapter_api

import (
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"todo-app-wbff/configs"
	corePort "todo-app-wbff/internal/pkg/application/core/port"
)

func RegisterAPI(
	e *echo.Echo,
	environmentVariableAccessor corePort.EnvironmentVariableAccessor,
	statusHandler StatusHandler,
	tokenClaimHandler TokenClaimHandler,
	todoHandler TodoHandler,
) {
	jwtConfig := echoJWT.Config{
		SigningKey:  []byte(environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameTokenSigningKey)),
		TokenLookup: "header:Authorization:Bearer ",
	}
	e.GET("/status", statusHandler.Get)
	auth := e.Group("/auth")
	auth.POST("/token-claims", tokenClaimHandler.Post)
	api := e.Group("/api")
	api.Use(echoJWT.WithConfig(jwtConfig))
	api.GET("/todos", todoHandler.GetCollection)
	api.POST("/todos", todoHandler.Post)
	api.PATCH("/todos/:todo", todoHandler.Patch)
	api.DELETE("/todos/:todo", todoHandler.Delete)
}
