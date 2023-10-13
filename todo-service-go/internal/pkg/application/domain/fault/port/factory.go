package domain_fault_port

type Factory interface {
	Create(cause error, code string, message string) Fault
	WrapError(err error, fallbackCode string) Fault
}
