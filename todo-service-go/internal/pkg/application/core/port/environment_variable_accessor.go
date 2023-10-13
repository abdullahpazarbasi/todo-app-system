package core_port

import "context"

type EnvironmentVariableAccessorKey struct{}

type EnvironmentVariableAccessor interface {
	NewContextWith(parentContext context.Context) context.Context
	Get(key string, defaultValue string) string
	GetOrThrowError(key string) (string, error)
	GetOrPanic(key string) string
}
