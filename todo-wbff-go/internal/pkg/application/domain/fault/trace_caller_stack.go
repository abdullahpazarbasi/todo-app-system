package domain_fault

import (
	"runtime"
	corePort "todo-app-wbff/internal/pkg/application/domain/fault/port"
)

func traceCallerStack(numberOfSkippableFrames int, depth int) *[]corePort.CallerFrame {
	callerProgramCounters := make([]uintptr, depth)
	n := runtime.Callers(numberOfSkippableFrames, callerProgramCounters)
	stack := make([]corePort.CallerFrame, 0)
	if n > 0 {
		var i int
		var k int
		var callPointProgramCounter uintptr
		var callerEntryPointProgramCounter uintptr
		var callPointFunc *runtime.Func
		var entryPointFunc *runtime.Func
		var callerFilePath string
		var callPointLine int
		var callerEntryPointLine int
		for i, k = 0, n; i < n; i++ {
			k--
			callPointProgramCounter = callerProgramCounters[i]
			callPointFunc = runtime.FuncForPC(callPointProgramCounter)
			callerEntryPointProgramCounter = callPointFunc.Entry()
			entryPointFunc = runtime.FuncForPC(callerEntryPointProgramCounter)
			callerFilePath, callPointLine = callPointFunc.FileLine(callPointProgramCounter)
			_, callerEntryPointLine = entryPointFunc.FileLine(callerEntryPointProgramCounter)
			stack = append(
				stack,
				&callerFrame{
					stackIndex:           k,
					callerFilePath:       callerFilePath,
					callPointLine:        callPointLine,
					callerEntryPointLine: callerEntryPointLine,
					callerName:           callPointFunc.Name(),
				},
			)
		}
	}

	return &stack
}
