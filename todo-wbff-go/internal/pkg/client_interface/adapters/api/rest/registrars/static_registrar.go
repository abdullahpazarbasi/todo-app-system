package client_interface_rest_api_registrars

import (
	"github.com/labstack/echo/v4"
	restApiHandlers "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/handlers"
)

func RegisterStaticAPI(e *echo.Echo, h restApiHandlers.HelloHandler) error {
	e.GET("/", h.Get)
	e.File("/favicon.ico", "web/static/favicon.ico")
	e.Static("/", "web/static")

	return nil
}
