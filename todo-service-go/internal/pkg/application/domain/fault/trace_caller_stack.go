package domain_fault

import (
	"runtime"
	"todo-app-service/internal/pkg/application/core/port"
)

func traceCallerStack(numberOfSkippableFrames int, depth int) *[]core_port.CallerFrame {
	callerProgramCounters := make([]uintptr, depth)
	n := runtime.Callers(numberOfSkippableFrames, callerProgramCounters)
	stack := make([]core_port.CallerFrame, 0)
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
					stackIndex:                     k,
					callerFilePath:                 callerFilePath,
					callerName:                     callPointFunc.Name(),
					callerEntryPointProgramCounter: callerEntryPointProgramCounter,
					callerEntryPointLine:           callerEntryPointLine,
					callPointProgramCounter:        callPointProgramCounter,
					callPointLine:                  callPointLine,
				},
			)
		}
	}

	return &stack
}
