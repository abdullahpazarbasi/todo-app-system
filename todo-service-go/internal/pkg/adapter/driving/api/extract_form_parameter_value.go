package driving_adapter_api

import (
	"fmt"
	"net/url"
)

func extractFormParameterValue(form url.Values, parameterName string, required bool) (string, error) {
	v, e := form[parameterName]
	if !e {
		if required {
			return "", fmt.Errorf("%s must be given", parameterName)
		} else {
			return "", nil
		}
	}
	if len(v) != 1 {
		return "", fmt.Errorf("multiple %s values are confusing", parameterName)
	}

	return v[0], nil
}
