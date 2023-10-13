package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHelloHandler(t *testing.T) {
	// when
	h := NewHelloHandler()

	// then
	assert.IsType(t, &helloHandler{}, h)
}

func TestHelloHandler_Get(t *testing.T) {
	// Setup
	e := echo.New()
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rc := httptest.NewRecorder()
	ec := e.NewContext(rq, rc)
	ec.SetPath("/")
	h := NewHelloHandler()

	// Assertions
	if assert.NoError(t, h.Get(ec)) {
		assert.Equal(t, http.StatusOK, rc.Code)
		assert.Equal(t, "Hello, World!", rc.Body.String())
	}
}
