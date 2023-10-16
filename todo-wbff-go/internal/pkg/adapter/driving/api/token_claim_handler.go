package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	drivingAdapterApiViews "todo-app-wbff/internal/pkg/adapter/driving/api/views"
	usecasePort "todo-app-wbff/internal/pkg/application/usecase/port"
)

type TokenClaimHandler interface {
	Post(ec echo.Context) error
}

type tokenClaimHandler struct {
	authService usecasePort.AuthService
}

func NewTokenClaimHandler(
	authService usecasePort.AuthService,
) TokenClaimHandler {
	return &tokenClaimHandler{
		authService: authService,
	}
}

func (h *tokenClaimHandler) Post(ec echo.Context) error {
	username := ec.FormValue("username")
	password := ec.FormValue("password")
	tokenEncoded, flt := h.authService.ClaimTokenByCredentials(username, password)
	if flt != nil {
		return flt
	}

	rs := drivingAdapterApiViews.TokenClaim{
		Token: tokenEncoded,
	}

	return ec.JSON(http.StatusCreated, &rs)
}
