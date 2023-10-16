package domain_fault_port

type FaultProposedHTTPStatusCodeOption interface {
	FaultOption
	ProposedHTTPStatusCode() int
}
