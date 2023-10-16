package core_port

type EnvironmentVariableAccessor interface {
	Get(key string, defaultValue string) string
	GetOrThrowError(key string) (string, error)
	GetOrPanic(key string) string
}
