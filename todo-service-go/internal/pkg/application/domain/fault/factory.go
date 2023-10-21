package domain_fault

import (
	"net/http"
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
		flt.traceCallerStack(numberOfSkippableFrames, depth)
	}

	return flt
}

func (f *factory) WrapError(err error, options ...domainFaultPort.FaultOption) domainFaultPort.Fault {
	switch e := err.(type) {
	case domainFaultPort.Fault:
		return e
	default:
		options = append(options, f.Cause(err))
		flt := f.createFault(&options)
		if f.isDebugModeOn() {
			flt.traceCallerStack(numberOfSkippableFrames, depth)
		}

		return flt
	}
}

func (f *factory) DenormalizeError(normalized map[string]interface{}, options ...domainFaultPort.FaultOption) domainFaultPort.Fault {
	return f.denormalizeError(normalized, &options)
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
		var proposedHTTPStatusCodeRaw float64
		proposedHTTPStatusCodeRaw, fit = proposedHTTPStatusCodeCandidate.(float64)
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
		var traceRaw []interface{}
		traceRaw, fit = traceCandidate.([]interface{})
		if !fit {
			panic("malformed trace")
		}
		var currentCallerFrame callerFrame
		var stackIndexCandidate interface{}
		var callerFilePathCandidate interface{}
		var callPointLineCandidate interface{}
		var callerEntryPointLineCandidate interface{}
		var callerNameCandidate interface{}
		var stackIndexRaw float64
		var callerFilePath string
		var callPointLineRaw float64
		var callerEntryPointLineRaw float64
		var callerName string
		for _, itemRaw := range traceRaw {
			var item map[string]interface{}
			item, fit = itemRaw.(map[string]interface{})
			if !fit {
				panic("malformed trace item")
			}

			currentCallerFrame = callerFrame{}

			stackIndexCandidate, existent = item["stack_index"]
			if existent {
				stackIndexRaw, fit = stackIndexCandidate.(float64)
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
				callPointLineRaw, fit = callPointLineCandidate.(float64)
				if !fit {
					panic("malformed call point line")
				}
				currentCallerFrame.callPointLine = int(callPointLineRaw)
			}

			callerEntryPointLineCandidate, existent = item["caller_entry_point_line"]
			if existent {
				callerEntryPointLineRaw, fit = callerEntryPointLineCandidate.(float64)
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

			*flt.callerFrames = append(*flt.callerFrames, &currentCallerFrame)
		}
	}

	return flt
}

func (f *factory) createFault(options *[]domainFaultPort.FaultOption) *fault {
	var tipe domainFaultPort.FaultType
	tipe, options = extractTypeFromFaultOptions(options)
	var message string
	message, options = extractMessageFromFaultOptions(options, "an error occurred")
	var proposedHTTPStatusCode int
	proposedHTTPStatusCode, options = extractProposedHTTPStatusCodeFromFaultOptions(options, http.StatusInternalServerError)
	var code string
	code, options = extractCodeFromFaultOptions(options, "")
	var cause error
	cause, options = extractCauseFromFaultOptions(options)

	callerFrames := make([]domainFaultPort.CallerFrame, 0)
	flt := &fault{
		tipe:                   tipe,
		code:                   code,
		proposedHTTPStatusCode: proposedHTTPStatusCode,
		message:                message,
		cause:                  cause,
		callerFrames:           &callerFrames,
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
