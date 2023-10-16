package core_port

type EnvironmentVariableAccessor interface {
	Get(name string, defaultValue string) string
	GetOrThrowError(name string) (string, error)
	GetOrPanic(name string) string
}
