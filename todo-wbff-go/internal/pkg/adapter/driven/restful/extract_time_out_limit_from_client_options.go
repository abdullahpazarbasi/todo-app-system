package driven_adapter_restful

import (
	"time"
)

func extractTimeOutLimitFromClientOptions(
	options *[]ClientOption,
) (
	time.Duration,
	*[]ClientOption,
) {
	var limit time.Duration
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case TimeOutLimitOption:
			limit = o.Limit()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return limit, &remainingOptions
}
