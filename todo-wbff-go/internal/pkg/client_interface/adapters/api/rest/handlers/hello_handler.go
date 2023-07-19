package client_interface_adapters_rest_api_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloHandler interface {
	Get(ec echo.Context) error
}

type helloHandler struct{}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (e *helloHandler) Get(ec echo.Context) error {
	return ec.String(http.StatusOK, "Hello, World!")
}
