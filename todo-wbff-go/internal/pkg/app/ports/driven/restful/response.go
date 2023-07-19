package driven_app_ports_restful

import (
	"io"
	"time"
)

type Response interface {
	Proto() string
	StatusCode() int
	Status() string
	IsStatusSuccess() bool
	IsStatusError() bool
	Header() *map[string][]string
	Cookies() *[]Cookie
	RawBody() io.ReadCloser
	Body() []byte
	String() string
	Size() int64
	ReceivedAt() *time.Time
	DecodeModel() (map[string]interface{}, error)
	DecodeCollection() ([]map[string]interface{}, error)
}
