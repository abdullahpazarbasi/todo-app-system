package driving_adapter_api_middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
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
