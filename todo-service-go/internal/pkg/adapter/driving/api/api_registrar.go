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
	api.GET("/users/:id/todos", todoHandler.GetCollectionOfUser)
	api.PUT("/todos/:id", todoHandler.Put)
	api.DELETE("/todos/:id", todoHandler.Delete)
}
