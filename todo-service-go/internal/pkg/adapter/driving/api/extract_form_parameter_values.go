package driving_adapter_api

import (
	"fmt"
	"net/url"
)

func extractFormParameterValues(form url.Values, parameterName string, required bool) ([]string, error) {
	var err error
	v, e := form[parameterName]
	if !e {
		if required {
			err = fmt.Errorf("%s must be given", parameterName)
		}
	}

	return v, err
}
