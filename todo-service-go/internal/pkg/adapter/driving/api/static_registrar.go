package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
)

func RegisterStaticAPI(e *echo.Echo, helloHandler HelloHandler) {
	e.GET("/", helloHandler.Get)
	e.File("/favicon.ico", "web/static/favicon.ico")
	e.Static("/", "web/static")
}
