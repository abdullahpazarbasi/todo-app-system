package domain_fault

import (
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
)

const numberOfSkippableFrames int = 3
const depth int = 32

type factory struct {
}

func NewFactory() domainFaultPort.Factory {
	return &factory{}
}

func (f *factory) Create(cause error, code string, message string) domainFaultPort.Fault {
	return &fault{
		tipe:                   "",
		code:                   code,
		proposedHTTPStatusCode: 0,
		message:                message,
		cause:                  cause,
		callerFrames:           traceCallerStack(numberOfSkippableFrames, depth),
	}
}

func (f *factory) WrapError(err error, fallbackCode string) domainFaultPort.Fault {
	switch e := err.(type) {
	case domainFaultPort.Fault:
		return e
	default:
		return &fault{
			code:                   fallbackCode,
			proposedHTTPStatusCode: 0,
			message:                "",
			cause:                  e,
			callerFrames:           traceCallerStack(numberOfSkippableFrames, depth),
		}
	}
}
