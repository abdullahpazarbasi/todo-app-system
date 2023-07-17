package driving_app_ports_error

type HttpError interface {
	Code() int
	IsClient() bool
	IsServer() bool
	DoesCodeMatchAnyOf(codes ...int) bool
	Error() string
	String() string
	Cause() error
}
