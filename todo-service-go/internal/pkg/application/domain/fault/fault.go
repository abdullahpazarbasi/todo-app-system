package domain_fault

import (
	"fmt"
	"net/http"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
)

type fault struct {
	tipe                   domainFaultPort.FaultType
	code                   string
	proposedHTTPStatusCode int
	message                string
	cause                  error
	callerFrames           *[]domainFaultPort.CallerFrame
}

func (e *fault) Type() domainFaultPort.FaultType {
	if e.tipe == "" {
		return domainFaultPort.FaultTypeUnknown
	}

	return e.tipe
}

func (e *fault) IsInType(typeCandidate domainFaultPort.FaultType) bool {
	return e.tipe == typeCandidate
}

func (e *fault) Code() string {
	return e.code
}

func (e *fault) ProposedHTTPStatusCode() int {
	return e.proposedHTTPStatusCode
}

func (e *fault) IsClient() bool {
	return e.proposedHTTPStatusCode > 399 && e.proposedHTTPStatusCode < 500
}

func (e *fault) IsServer() bool {
	return e.proposedHTTPStatusCode > 499
}

func (e *fault) DoesProposedHTTPStatusCodeMatchAnyOf(codes ...int) bool {
	for _, code := range codes {
		if e.proposedHTTPStatusCode == code {
			return true
		}
	}

	return false
}

func (e *fault) Error() string {
	return e.String()
}

func (e *fault) String() string {
	if e.cause == nil {
		return fmt.Sprintf(
			"code=%s, status=%d, message=%v, trace=%v",
			e.code,
			e.proposedHTTPStatusCode,
			e.message,
			e.printCallerFrames(),
		)
	}

	return fmt.Sprintf(
		"code=%s, status=%d, message=%v, trace=%v, cause(%v)",
		e.code,
		e.proposedHTTPStatusCode,
		e.message,
		e.printCallerFrames(),
		e.cause,
	)
}

func (e *fault) Cause() error {
	return e.cause
}

func (e *fault) Normalize(fullDetailed bool) map[string]interface{} {
	err := make(map[string]interface{})
	err["code"] = e.code
	err["proposed_http_status_code"] = e.proposedHTTPStatusCode
	if fullDetailed {
		err["message"] = e.message
		callerFrames := make([]*map[string]interface{}, 0)
		for _, cf := range *e.callerFrames {
			callerFrames = append(callerFrames, &map[string]interface{}{
				"stack_index":             cf.StackIndex(),
				"caller_file_path":        cf.CallerFilePath(),
				"call_point_line":         cf.CallPointLine(),
				"caller_entry_point_line": cf.CallerEntryPointLine(),
				"caller_name":             cf.CallerName(),
			})
		}
		err["trace"] = callerFrames
		if e.cause != nil {
			var cause map[string]interface{}
			switch c := e.cause.(type) {
			case domainFaultPort.Fault:
				cause = c.Normalize(fullDetailed)
			default:
				cause = map[string]interface{}{"message": c.Error()}
			}
			err["cause"] = &cause
		}
	} else {
		if e.proposedHTTPStatusCode > 99 {
			err["message"] = http.StatusText(e.proposedHTTPStatusCode)
		} else {
			err["message"] = "an error occurred"
		}
	}

	return err
}

func (e *fault) CallerFrames() *[]domainFaultPort.CallerFrame {
	return e.callerFrames
}

func (e *fault) printCallerFrames() string {
	var o string
	for i, cf := range *e.callerFrames {
		if i > 0 {
			o = fmt.Sprintf("%s, ", o)
		}
		o = fmt.Sprintf(
			"%si:%d crfp:%s cpl:%d crn:%s crepl:%d",
			o,
			cf.StackIndex(),
			cf.CallerFilePath(),
			cf.CallPointLine(),
			cf.CallerName(),
			cf.CallerEntryPointLine(),
		)
	}

	return o
}
