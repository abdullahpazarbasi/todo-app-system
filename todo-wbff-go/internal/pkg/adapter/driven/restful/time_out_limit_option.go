package driven_adapter_restful

import "time"

type TimeOutLimitOption interface {
	ClientOption
	Limit() time.Duration
}
