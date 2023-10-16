package driving_adapter_api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strconv"
	"todo-app-service/configs"
	drivingAdapterApiMiddlewares "todo-app-service/internal/pkg/adapter/driving/api/middlewares"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

func RegisterMiddlewares(
	e *echo.Echo,
	eva corePort.EnvironmentVariableAccessor,
	parentContext context.Context,
) {
	var err error
	e.Debug, err = strconv.ParseBool(eva.Get(configs.EnvironmentVariableNameAppDebug, "false"))
	if err != nil {
		log.Fatalf("malformed %s: %v", configs.EnvironmentVariableNameAppDebug, e)
	}
	e.HTTPErrorHandler = errorHandler
	e.Use(drivingAdapterApiMiddlewares.OverrideParentContext(parentContext))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
