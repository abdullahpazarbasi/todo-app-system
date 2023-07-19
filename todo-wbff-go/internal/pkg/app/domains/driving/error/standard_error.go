package driving_app_domains_error

import (
	"fmt"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

type serviceError struct {
	code                   string
	proposedHTTPStatusCode int
	message                string
	cause                  error
}

func NewServiceError(cause error, code string, proposedHTTPStatusCode int, message string) *serviceError {
	return &serviceError{
		code:                   code,
		proposedHTTPStatusCode: proposedHTTPStatusCode,
		message:                message,
		cause:                  cause,
	}
}

func (e *serviceError) Code() string {
	return e.code
}

func (e *serviceError) ProposedHTTPStatusCode() int {
	return e.proposedHTTPStatusCode
}

func (e *serviceError) IsClient() bool {
	return e.proposedHTTPStatusCode > 399 && e.proposedHTTPStatusCode < 500
}

func (e *serviceError) IsServer() bool {
	return e.proposedHTTPStatusCode > 499
}

func (e *serviceError) DoesProposedHTTPStatusCodeMatchAnyOf(codes ...int) bool {
	for _, code := range codes {
		if e.proposedHTTPStatusCode == code {
			return true
		}
	}

	return false
}

func (e *serviceError) Error() string {
	return e.String()
}

func (e *serviceError) String() string {
	if e.cause == nil {
		return fmt.Sprintf(
			"code=%s, status=%d, message=%v",
			e.code,
			e.proposedHTTPStatusCode,
			e.message,
		)
	}

	return fmt.Sprintf(
		"code=%s, status=%d, message=%v, cause:%v:",
		e.code,
		e.proposedHTTPStatusCode,
		e.message,
		e.cause,
	)
}

func (e *serviceError) Cause() error {
	return e.cause
}

func (e *serviceError) Normalize(fullDetailed bool) map[string]interface{} {
	return normalize(e, fullDetailed)
}

func normalize(err error, fullDetailed bool) map[string]interface{} {
	normalized := make(map[string]interface{})

	var c error
	switch e := err.(type) {
	case *serviceError:
		if e.code != "" {
			normalized["code"] = e.code
		}
		if e.message != "" {
			normalized["message"] = e.message
		}
		c = e.cause
	case drivenAppPortsRestful.Error:
		code := e.Code()
		if code != "" {
			normalized["code"] = code
		}
		message := e.Error()
		if message != "" {
			normalized["message"] = message
		}
		c = e.Cause()
	default:
		normalized["message"] = "an error occurred"
	}
	if fullDetailed {
		if c != nil {
			normalized["cause"] = normalize(c, fullDetailed)
		}
	}

	return normalized
}
