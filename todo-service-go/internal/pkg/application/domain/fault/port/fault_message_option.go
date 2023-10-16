package domain_fault_port

type FaultMessageOption interface {
	FaultOption
	Message() string
}
