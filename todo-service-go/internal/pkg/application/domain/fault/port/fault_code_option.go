package domain_fault_port

type FaultCodeOption interface {
	FaultOption
	Code() string
}
