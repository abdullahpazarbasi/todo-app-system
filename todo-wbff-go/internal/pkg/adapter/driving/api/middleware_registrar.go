package driving_adapter_api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"todo-app-wbff/configs"
	drivingAdapterApiMiddlewares "todo-app-wbff/internal/pkg/adapter/driving/api/middlewares"
	corePort "todo-app-wbff/internal/pkg/application/core/port"
)

func RegisterMiddlewares(
	e *echo.Echo,
	environmentVariableAccessor corePort.EnvironmentVariableAccessor,
	parentContext context.Context,
) {
	var err error
	e.Debug, err = strconv.ParseBool(environmentVariableAccessor.Get(configs.EnvironmentVariableNameAppDebug, "false"))
	if err != nil {
		log.Fatalf("malformed %s: %v", configs.EnvironmentVariableNameAppDebug, e)
	}
	e.HTTPErrorHandler = errorHandler
	e.Use(drivingAdapterApiMiddlewares.OverrideParentContext(parentContext))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
	}))
}
