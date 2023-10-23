package driven_adapter_restful

import (
	"time"
)

func extractTimeOutLimitFromClientOptions(
	options *[]ClientOption,
	defaultValue time.Duration,
) (
	time.Duration,
	*[]ClientOption,
) {
	var existent bool
	var limit time.Duration
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case TimeOutLimitOption:
			limit = o.Limit()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if existent {
		limit = defaultValue
	}

	return limit, &remainingOptions
}
