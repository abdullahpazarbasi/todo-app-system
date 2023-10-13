package domain_fault_port

import "todo-app-service/internal/pkg/application/core/port"

type Fault interface {
	Type() FaultType
	Code() string
	ProposedHTTPStatusCode() int
	IsClient() bool
	IsServer() bool
	DoesProposedHTTPStatusCodeMatchAnyOf(codes ...int) bool
	Error() string
	String() string
	Cause() error
	Normalize(fullDetailed bool) map[string]interface{}
	CallerFrames() *[]core_port.CallerFrame
	SetType(tipe FaultType) Fault
	SetProposedHTTPStatusCode(proposedHTTPStatusCode int) Fault
	SetMessage(message string) Fault
}
