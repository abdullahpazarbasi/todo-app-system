package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type StatusHandler interface {
	Get(ec echo.Context) error
}

type statusHandler struct{}

func NewStatusHandler() StatusHandler {
	return &statusHandler{}
}

func (e *statusHandler) Get(ec echo.Context) error {
	return ec.String(http.StatusOK, "OK")
}
