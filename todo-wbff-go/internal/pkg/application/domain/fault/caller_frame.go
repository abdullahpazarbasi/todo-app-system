package domain_fault

type callerFrame struct {
	stackIndex           int
	callerFilePath       string
	callPointLine        int
	callerEntryPointLine int
	callerName           string
}

func (f *callerFrame) StackIndex() int {
	return f.stackIndex
}

func (f *callerFrame) CallerFilePath() string {
	return f.callerFilePath
}

func (f *callerFrame) CallPointLine() int {
	return f.callPointLine
}

func (f *callerFrame) CallerEntryPointLine() int {
	return f.callerEntryPointLine
}

func (f *callerFrame) CallerName() string {
	return f.callerName
}
