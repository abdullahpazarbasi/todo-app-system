package driven_app_ports_restful

import "time"

type Request interface {
	SentAt() *time.Time
}
