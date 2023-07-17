package client_interface_rest_api_middlewares

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type ParentContextKey struct{}

func OverrideParentContext(parentContext context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			rq := ec.Request()
			ec.SetRequest(rq.WithContext(parentContext))

			return next(ec)
		}
	}
}
