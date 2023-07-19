package driving_app_ports_error

type ServiceError interface {
	Code() string
	ProposedHTTPStatusCode() int
	IsClient() bool
	IsServer() bool
	DoesProposedHTTPStatusCodeMatchAnyOf(codes ...int) bool
	Error() string
	String() string
	Cause() error
	Normalize(fullDetailed bool) map[string]interface{}
}
