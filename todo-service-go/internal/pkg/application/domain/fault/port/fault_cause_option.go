package domain_fault_port

type FaultCauseOption interface {
	FaultOption
	Cause() error
}
