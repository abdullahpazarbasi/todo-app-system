package domain_fault

import (
	"fmt"
	"net/http"
	"todo-app-service/internal/pkg/application/core/port"
	"todo-app-service/internal/pkg/application/domain/fault/port"
)

type fault struct {
	tipe                   domain_fault_port.FaultType
	code                   string
	proposedHTTPStatusCode int
	message                string
	cause                  error
	callerFrames           *[]core_port.CallerFrame
}

func (e *fault) Type() domain_fault_port.FaultType {
	if e.tipe == "" {
		return domain_fault_port.FaultTypeUnknown
	}

	return e.tipe
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
		if e.cause != nil {
			var cause map[string]interface{}
			switch c := e.cause.(type) {
			case domain_fault_port.Fault:
				cause = c.Normalize(fullDetailed)
			default:
				cause = map[string]interface{}{"message": c.Error()}
			}
			err["cause"] = &cause
		}
		callerFrames := make([]*map[string]interface{}, 0)
		for _, cf := range *e.callerFrames {
			callerFrames = append(callerFrames, &map[string]interface{}{
				"stack_index":                        cf.StackIndex(),
				"caller_file_path":                   cf.CallerFilePath(),
				"caller_name":                        cf.CallerName(),
				"caller_entry_point_program_counter": cf.CallerEntryPointProgramCounter(),
				"caller_entry_point_line":            cf.CallerEntryPointLine(),
				"call_point_program_counter":         cf.CallPointProgramCounter(),
				"call_point_line":                    cf.CallPointLine(),
			})
		}
		err["trace"] = callerFrames
	} else {
		if e.proposedHTTPStatusCode > 99 {
			err["message"] = http.StatusText(e.proposedHTTPStatusCode)
		} else {
			err["message"] = "an error occurred"
		}
	}

	return err
}

func (e *fault) CallerFrames() *[]core_port.CallerFrame {
	return e.callerFrames
}

func (e *fault) SetType(tipe domain_fault_port.FaultType) domain_fault_port.Fault {
	if e.tipe != "" {
		panic("fault type is already set")
	}
	e.tipe = tipe

	return e
}

func (e *fault) SetProposedHTTPStatusCode(proposedHTTPStatusCode int) domain_fault_port.Fault {
	if e.proposedHTTPStatusCode != 0 {
		panic("proposed HTTP status code is already set")
	}
	e.proposedHTTPStatusCode = proposedHTTPStatusCode

	return e
}

func (e *fault) SetMessage(message string) domain_fault_port.Fault {
	if e.message != "" {
		panic("message is already set")
	}
	e.message = message

	return e
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
