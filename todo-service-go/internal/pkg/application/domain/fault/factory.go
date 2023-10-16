package domain_fault

import (
	"strconv"
	corePort "todo-app-service/internal/pkg/application/core/port"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
)

const numberOfSkippableFrames int = 3
const depth int = 32

type factory struct {
	environmentVariableAccessor      corePort.EnvironmentVariableAccessor
	debugModeEnvironmentVariableName string
}

func NewFactory(
	environmentVariableAccessor corePort.EnvironmentVariableAccessor,
	debugModeEnvironmentVariableName string,
) domainFaultPort.Factory {
	return &factory{
		environmentVariableAccessor:      environmentVariableAccessor,
		debugModeEnvironmentVariableName: debugModeEnvironmentVariableName,
	}
}

func (f *factory) CreateFault(options ...domainFaultPort.FaultOption) domainFaultPort.Fault {
	flt := f.createFault(&options)
	if f.isDebugModeOn() {
		flt.callerFrames = traceCallerStack(numberOfSkippableFrames, depth)
	}

	return flt
}

func (f *factory) WrapError(err error, options ...domainFaultPort.FaultOption) domainFaultPort.Fault {
	switch e := err.(type) {
	case domainFaultPort.Fault:
		return e
	default:
		options = append(options, f.Cause(err))

		return f.createFault(&options)
	}
}

func (f *factory) DenormalizeError(normalized *map[string]interface{}, options ...domainFaultPort.FaultOption) domainFaultPort.Fault {
	return f.denormalizeError(*normalized, &options)
}

func (f *factory) Cause(cause error) domainFaultPort.FaultCauseOption {
	return &faultCauseOption{
		cause: cause,
	}
}

func (f *factory) Type(tipe domainFaultPort.FaultType) domainFaultPort.FaultTypeOption {
	return &faultTypeOption{
		tipe: tipe,
	}
}

func (f *factory) Code(code string) domainFaultPort.FaultCodeOption {
	return &faultCodeOption{
		code: code,
	}
}

func (f *factory) ProposedHTTPStatusCode(proposedHTTPStatusCode int) domainFaultPort.FaultProposedHTTPStatusCodeOption {
	return &faultProposedHTTPStatusCodeOption{
		proposedHTTPStatusCode: proposedHTTPStatusCode,
	}
}

func (f *factory) Message(message string) domainFaultPort.FaultMessageOption {
	return &faultMessageOption{
		message: message,
	}
}

func (f *factory) denormalizeError(
	normalized map[string]interface{},
	options *[]domainFaultPort.FaultOption,
) domainFaultPort.Fault {
	var existent, fit bool

	var typeCandidate interface{}
	typeCandidate, existent = normalized["type"]
	if existent {
		var tipe string
		tipe, fit = typeCandidate.(string)
		if !fit {
			panic("malformed HTTP error type")
		}
		*options = append(*options, f.Type(translateFaultTypeFromString(tipe)))
	}

	var codeCandidate interface{}
	codeCandidate, existent = normalized["code"]
	if existent {
		var code string
		code, fit = codeCandidate.(string)
		if !fit {
			panic("malformed HTTP error code")
		}
		*options = append(*options, f.Code(code))
	}

	var messageCandidate interface{}
	messageCandidate, existent = normalized["message"]
	if existent {
		var message string
		message, fit = messageCandidate.(string)
		if !fit {
			panic("malformed HTTP error message")
		}
		*options = append(*options, f.Message(message))
	}

	var proposedHTTPStatusCodeCandidate interface{}
	proposedHTTPStatusCodeCandidate, existent = normalized["proposed_http_status_code"]
	if existent {
		var proposedHTTPStatusCodeRaw int64
		proposedHTTPStatusCodeRaw, fit = proposedHTTPStatusCodeCandidate.(int64)
		if !fit {
			panic("malformed HTTP status code")
		}
		*options = append(*options, f.ProposedHTTPStatusCode(int(proposedHTTPStatusCodeRaw)))
	}

	var causeCandidate interface{}
	causeCandidate, existent = normalized["cause"]
	if existent {
		var subCause map[string]interface{}
		subCause, fit = causeCandidate.(map[string]interface{})
		if !fit {
			panic("malformed HTTP error cause")
		}
		subOptions := make([]domainFaultPort.FaultOption, 0)
		*options = append(*options, f.Cause(f.denormalizeError(subCause, &subOptions)))
	}

	flt := f.createFault(options)

	var traceCandidate interface{}
	traceCandidate, existent = normalized["trace"]
	if existent {
		var traceRaw []map[string]interface{}
		traceRaw, fit = traceCandidate.([]map[string]interface{})
		if !fit {
			panic("malformed trace")
		}
		callerFrames := make([]domainFaultPort.CallerFrame, 0)
		var currentCallerFrame callerFrame
		var stackIndexCandidate interface{}
		var callerFilePathCandidate interface{}
		var callPointLineCandidate interface{}
		var callerEntryPointLineCandidate interface{}
		var callerNameCandidate interface{}
		var stackIndexRaw int64
		var callerFilePath string
		var callPointLineRaw int64
		var callerEntryPointLineRaw int64
		var callerName string
		for _, item := range traceRaw {
			currentCallerFrame = callerFrame{}

			stackIndexCandidate, existent = item["stack_index"]
			if existent {
				stackIndexRaw, fit = stackIndexCandidate.(int64)
				if !fit {
					panic("malformed stack index")
				}
				currentCallerFrame.stackIndex = int(stackIndexRaw)
			}

			callerFilePathCandidate, existent = item["caller_file_path"]
			if existent {
				callerFilePath, fit = callerFilePathCandidate.(string)
				if !fit {
					panic("malformed caller file path")
				}
				currentCallerFrame.callerFilePath = callerFilePath
			}

			callPointLineCandidate, existent = item["call_point_line"]
			if existent {
				callPointLineRaw, fit = callPointLineCandidate.(int64)
				if !fit {
					panic("malformed call point line")
				}
				currentCallerFrame.callPointLine = int(callPointLineRaw)
			}

			callerEntryPointLineCandidate, existent = item["caller_entry_point_line"]
			if existent {
				callerEntryPointLineRaw, fit = callerEntryPointLineCandidate.(int64)
				if !fit {
					panic("malformed caller entry point line")
				}
				currentCallerFrame.callerEntryPointLine = int(callerEntryPointLineRaw)
			}

			callerNameCandidate, existent = item["caller_name"]
			if existent {
				callerName, fit = callerNameCandidate.(string)
				if !fit {
					panic("malformed caller name")
				}
				currentCallerFrame.callerName = callerName
			}

			callerFrames = append(callerFrames, &currentCallerFrame)
		}
		flt.callerFrames = &callerFrames
	}

	return flt
}

func (f *factory) createFault(options *[]domainFaultPort.FaultOption) *fault {
	var tipe domainFaultPort.FaultType
	tipe, options = extractTypeFromFaultOptions(options)
	var message string
	message, options = extractMessageFromFaultOptions(options, "")
	var proposedHTTPStatusCode int
	proposedHTTPStatusCode, options = extractProposedHTTPStatusCodeFromFaultOptions(options, 0)
	var code string
	code, options = extractCodeFromFaultOptions(options, "")
	var cause error
	cause, options = extractCauseFromFaultOptions(options)

	flt := &fault{
		tipe:                   tipe,
		code:                   code,
		proposedHTTPStatusCode: proposedHTTPStatusCode,
		message:                message,
		cause:                  cause,
	}

	return flt
}

func (f *factory) isDebugModeOn() bool {
	debugRaw := f.environmentVariableAccessor.GetOrPanic(f.debugModeEnvironmentVariableName)
	debug, err := strconv.ParseBool(debugRaw)
	if err != nil {
		panic(err)
	}

	return debug
}
