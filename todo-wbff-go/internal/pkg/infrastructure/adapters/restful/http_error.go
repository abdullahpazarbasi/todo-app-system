package infrastructure_adapters_restful

type httpError struct {
	code           string
	httpStatusCode int
	message        string
	cause          error
}

func denormalize(source map[string]interface{}) *httpError {
	var existent, fit bool
	err := &httpError{}
	var codeCandidate interface{}
	codeCandidate, existent = source["code"]
	if existent {
		err.code, fit = codeCandidate.(string)
		if !fit {
			panic("malformed HTTP error code")
		}
	}
	var httpStatusCodeCandidate interface{}
	httpStatusCodeCandidate, existent = source["status_code"]
	if existent {
		err.httpStatusCode, fit = httpStatusCodeCandidate.(int)
		if !fit {
			panic("malformed HTTP error status code")
		}
	}
	var messageCandidate interface{}
	messageCandidate, existent = source["message"]
	if existent {
		err.message, fit = messageCandidate.(string)
		if !fit {
			panic("malformed HTTP error message")
		}
	}
	var causeCandidate interface{}
	causeCandidate, existent = source["cause"]
	if existent {
		var subSource map[string]interface{}
		subSource, fit = causeCandidate.(map[string]interface{})
		if !fit {
			panic("malformed HTTP error cause")
		}
		err.cause = denormalize(subSource)
	}

	return err
}

func (e *httpError) Code() string {
	return e.code
}

func (e *httpError) HTTPStatusCode() int {
	return e.httpStatusCode
}

func (e *httpError) IsClient() bool {
	return e.httpStatusCode > 399 && e.httpStatusCode < 500
}

func (e *httpError) IsServer() bool {
	return e.httpStatusCode > 499
}

func (e *httpError) DoesHTTPStatusCodeMatchAnyOf(codes ...int) bool {
	for _, code := range codes {
		if e.httpStatusCode == code {
			return true
		}
	}

	return false
}

func (e *httpError) Error() string {
	return e.message
}

func (e *httpError) String() string {
	return e.message
}

func (e *httpError) Cause() error {
	return e.cause
}
