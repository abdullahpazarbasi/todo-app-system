package client_interface_adapters_rest_api_handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHelloHandler(t *testing.T) {
	// when
	h := NewHelloHandler()

	// then
	assert.IsType(t, &helloHandler{}, h)
}
