package driven_app_ports_os

import "context"

type EnvironmentVariableAccessorKey struct{}

type EnvironmentVariableAccessor interface {
	NewContextWith(parentContext context.Context) context.Context
	Get(key string, defaultValue string) string
	GetOrPanic(key string) string
}
