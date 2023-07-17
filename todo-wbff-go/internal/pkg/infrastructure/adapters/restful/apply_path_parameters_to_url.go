package infrastructure_adapters_restful

import (
	"fmt"
	"net/url"
	"strings"
)

func applyPathParametersToURL(urlPattern string, pathParameters *map[string]string) string {
	newURL := urlPattern
	if len(*pathParameters) > 0 {
		for p, v := range *pathParameters {
			newURL = strings.Replace(newURL, fmt.Sprintf("{%s}", p), url.PathEscape(v), -1)
		}
	}

	return newURL
}
