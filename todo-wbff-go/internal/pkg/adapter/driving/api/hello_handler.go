package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloHandler interface {
	Get(ec echo.Context) error
}

type helloHandler struct{}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (e *helloHandler) Get(ec echo.Context) error {
	return ec.String(http.StatusOK, "Hello, World!")
}
