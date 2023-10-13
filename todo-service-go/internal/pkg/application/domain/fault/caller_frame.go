package domain_fault

type callerFrame struct {
	stackIndex                     int
	callerFilePath                 string
	callerName                     string
	callerEntryPointProgramCounter uintptr
	callerEntryPointLine           int
	callPointProgramCounter        uintptr
	callPointLine                  int
}

func (f *callerFrame) StackIndex() int {
	return f.stackIndex
}

func (f *callerFrame) CallerFilePath() string {
	return f.callerFilePath
}

func (f *callerFrame) CallerName() string {
	return f.callerName
}

func (f *callerFrame) CallerEntryPointProgramCounter() uintptr {
	return f.callerEntryPointProgramCounter
}

func (f *callerFrame) CallerEntryPointLine() int {
	return f.callerEntryPointLine
}

func (f *callerFrame) CallPointProgramCounter() uintptr {
	return f.callPointProgramCounter
}

func (f *callerFrame) CallPointLine() int {
	return f.callPointLine
}
