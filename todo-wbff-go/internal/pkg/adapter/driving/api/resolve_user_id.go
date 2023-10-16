package driving_adapter_api

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func resolveUserID(ec echo.Context) (string, error) {
	user := ec.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims.GetSubject()
}
