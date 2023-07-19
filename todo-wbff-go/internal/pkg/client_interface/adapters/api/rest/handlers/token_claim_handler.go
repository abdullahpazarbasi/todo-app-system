package client_interface_adapters_rest_api_handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	drivenAppPortsOs "todo-app-wbff/internal/pkg/app/domains/driven/os"
	drivenAppPortsCore "todo-app-wbff/internal/pkg/app/ports/driven/core"
	restApiViews "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/views"
)

type TokenClaimHandler interface {
	Post(ec echo.Context) error
}

type tokenClaimHandler struct{}

func NewTokenClaimHandler() *tokenClaimHandler {
	return &tokenClaimHandler{}
}

func (tr *tokenClaimHandler) Post(ec echo.Context) error {
	username := ec.FormValue("username")
	password := ec.FormValue("password")

	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	claims := &jwt.RegisteredClaims{
		Subject:   "1",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}

	eva := drivenAppPortsOs.ExtractEnvironmentVariableAccessorFromContext(ec.Request().Context())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenEncoded, err := token.SignedString([]byte(eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTokenSigningKey)))
	if err != nil {
		return err
	}

	rs := restApiViews.TokenClaim{
		Token: tokenEncoded,
	}

	return ec.JSON(http.StatusCreated, &rs)
}
