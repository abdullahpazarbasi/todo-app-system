package domain_fault_port

type FaultTypeOption interface {
	FaultOption
	Type() FaultType
}
