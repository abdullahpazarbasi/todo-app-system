package infrastructure_adapters_restful

import (
	"fmt"
	"net/url"
	appPortsRestfulModels "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

func applyQueryParametersToURL(targetURL string, queryParameters appPortsRestfulModels.ParameterMap) string {
	q := url.Values(queryParameters.Export())

	return fmt.Sprintf("%s?%s", targetURL, q.Encode())
}
