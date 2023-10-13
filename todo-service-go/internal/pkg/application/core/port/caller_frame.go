package core_port

type CallerFrame interface {
	StackIndex() int
	CallerFilePath() string
	CallerName() string
	CallerEntryPointProgramCounter() uintptr
	CallerEntryPointLine() int
	CallPointProgramCounter() uintptr
	CallPointLine() int
}
