package domain_fault_port

type Fault interface {
	Type() FaultType
	IsInType(typeCandidate FaultType) bool
	Code() string
	ProposedHTTPStatusCode() int
	IsClient() bool
	IsServer() bool
	DoesProposedHTTPStatusCodeMatchAnyOf(codes ...int) bool
	Error() string
	String() string
	Cause() error
	Normalize(fullDetailed bool) map[string]interface{}
	CallerFrames() *[]CallerFrame
}
