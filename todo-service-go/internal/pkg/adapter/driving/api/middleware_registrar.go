package driving_adapter_api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strconv"
	drivingAdapterApiMiddlewares "todo-app-service/internal/pkg/adapter/driving/api/middlewares"
	"todo-app-service/internal/pkg/application/core"
)

func RegisterMiddlewares(e *echo.Echo, parentContext context.Context) {
	var err error
	eva := core.ExtractEnvironmentVariableAccessorFromContext(parentContext)
	e.Debug, err = strconv.ParseBool(eva.Get("APP_DEBUG", "false"))
	if err != nil {
		log.Fatalf("malformed APP_DEBUG: %v", e)
	}
	e.HTTPErrorHandler = ErrorHandler
	e.Use(drivingAdapterApiMiddlewares.OverrideParentContext(parentContext))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
