package driven_app_ports_restful

import "time"

type TimeOutLimitOption interface {
	ClientOption
	Limit() time.Duration
}
