package driven_adapter_restful

import (
	"fmt"
	"net/url"
)

func applyQueryParametersToURL(targetURL string, queryParameters ParameterMap) string {
	if queryParameters == nil {
		return targetURL
	}
	q := url.Values(queryParameters.Export())

	return fmt.Sprintf("%s?%s", targetURL, q.Encode())
}
