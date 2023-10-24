package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
)

func RegisterAPIs(
	e *echo.Echo,
	statusHandler StatusHandler,
	todoHandler TodoHandler,
) {
	e.GET("/status", statusHandler.Get)
	api := e.Group("/api")
	api.POST("/todos", todoHandler.Post)
	api.GET("/users/:user/todos", todoHandler.GetCollectionOfUser)
	api.PATCH("/todos/:todo", todoHandler.Patch)
	api.DELETE("/todos/:todo", todoHandler.Delete)
}
