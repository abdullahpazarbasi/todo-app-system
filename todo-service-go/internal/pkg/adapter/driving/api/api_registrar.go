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
	api.PUT("/todos/:todo", todoHandler.Put)
	api.DELETE("/todos/:todo", todoHandler.Delete)
}
