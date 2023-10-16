package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
)

func RegisterStaticAPI(
	e *echo.Echo,
	helloHandler HelloHandler,
) {
	e.GET("/", helloHandler.Get)
	e.File("/favicon.ico", "internal/pkg/presentation/core/asset/icons/favicon.ico")
	e.Static("/", "internal/pkg/presentation/core/web/static")
}
