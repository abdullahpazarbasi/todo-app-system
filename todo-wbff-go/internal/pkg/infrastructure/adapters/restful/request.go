package infrastructure_adapters_restful

import (
	"net/http"
	"time"
)

type request struct {
	raw    *http.Request
	sentAt time.Time
}

func (r *request) SentAt() *time.Time {
	return &r.sentAt
}
