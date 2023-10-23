package domain_fault_port

type Factory interface {
	CreateFault(options ...FaultOption) Fault
	WrapError(err error, options ...FaultOption) Fault
	DenormalizeError(normalized map[string]interface{}, options ...FaultOption) Fault
	Cause(cause error) FaultCauseOption
	Type(tipe FaultType) FaultTypeOption
	Code(code string) FaultCodeOption
	ProposedHTTPStatusCode(proposedHTTPStatusCode int) FaultProposedHTTPStatusCodeOption
	Message(message string) FaultMessageOption
}
