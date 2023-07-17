package infrastructure_adapters_restful

import (
	"time"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

func extractTimeOutLimitFromClientOptions(
	options *[]drivenAppPortsRestful.ClientOption,
) (
	time.Duration,
	*[]drivenAppPortsRestful.ClientOption,
) {
	var limit time.Duration
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.TimeOutLimitOption:
			limit = o.Limit()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return limit, &remainingOptions
}
