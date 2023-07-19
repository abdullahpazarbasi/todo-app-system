package driven_app_ports_restful

type Error interface {
	Code() string
	HTTPStatusCode() int
	IsClient() bool
	IsServer() bool
	DoesHTTPStatusCodeMatchAnyOf(codes ...int) bool
	Error() string
	String() string
	Cause() error
}
