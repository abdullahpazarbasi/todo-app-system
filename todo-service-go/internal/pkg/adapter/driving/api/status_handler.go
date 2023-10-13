package driving_adapter_api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StatusHandler interface {
	Get(ec echo.Context) error
}

type statusHandler struct {
	databaseClient *sql.DB
}

func NewStatusHandler(databaseClient *sql.DB) StatusHandler {
	return &statusHandler{
		databaseClient: databaseClient,
	}
}

func (h *statusHandler) Get(ec echo.Context) error {
	err := h.databaseClient.PingContext(ec.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ec.String(http.StatusOK, "OK")
}
