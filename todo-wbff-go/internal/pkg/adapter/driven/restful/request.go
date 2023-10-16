package driven_adapter_restful

import (
	"net/http"
	"time"
)

type Request interface {
	SentAt() *time.Time
}

type request struct {
	raw    *http.Request
	sentAt time.Time
}

func (r *request) SentAt() *time.Time {
	return &r.sentAt
}
