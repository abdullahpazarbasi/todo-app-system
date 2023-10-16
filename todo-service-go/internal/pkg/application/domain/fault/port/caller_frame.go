package domain_fault_port

type CallerFrame interface {
	StackIndex() int
	CallerFilePath() string
	CallPointLine() int
	CallerEntryPointLine() int
	CallerName() string
}
