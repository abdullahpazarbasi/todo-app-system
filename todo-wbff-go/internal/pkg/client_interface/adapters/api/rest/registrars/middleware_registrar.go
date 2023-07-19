package client_interface_rest_api_registrars

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	clientInterfaceAdaptersRestApiHandlers "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/handlers"
	clientInterfaceAdaptersRestApiMiddlewares "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/middlewares"
)

func RegisterMiddlewares(e *echo.Echo, parentContext context.Context) error {
	e.HTTPErrorHandler = clientInterfaceAdaptersRestApiHandlers.ErrorHandler
	e.Use(clientInterfaceAdaptersRestApiMiddlewares.OverrideParentContext(parentContext))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
	}))

	return nil
}
