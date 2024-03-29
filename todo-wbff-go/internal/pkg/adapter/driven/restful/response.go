package driven_adapter_restful

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response interface {
	Proto() string
	StatusCode() int
	Status() string
	IsStatusSuccess() bool
	IsStatusError() bool
	Header() map[string][]string
	Cookies() *[]Cookie
	RawBody() io.ReadCloser
	Body() []byte
	String() string
	Size() int64
	ReceivedAt() *time.Time
	DecodeModel() (map[string]interface{}, error)
	DecodeCollection() ([]map[string]interface{}, error)
}

type response struct {
	raw        *http.Response
	receivedAt time.Time
}

func (r *response) Proto() string {
	if r.raw == nil {
		return ""
	}

	return r.raw.Proto
}

func (r *response) StatusCode() int {
	if r.raw == nil {
		return 0
	}

	return r.raw.StatusCode
}

func (r *response) Status() string {
	if r.raw == nil {
		return ""
	}

	return r.raw.Status
}

func (r *response) IsStatusSuccess() bool {
	if r.raw == nil {
		return false
	}

	return r.raw.StatusCode > 199 && r.raw.StatusCode < 300
}

func (r *response) IsStatusError() bool {
	if r.raw == nil {
		return false
	}

	return r.raw.StatusCode > 399
}

func (r *response) Header() map[string][]string {
	if r.raw == nil {
		return nil
	}

	return r.raw.Header
}

func (r *response) Cookies() *[]Cookie {
	cs := make([]Cookie, 0)
	if r.raw == nil {
		return &cs
	}

	for _, c := range r.raw.Cookies() {
		cs = append(cs, NewCookieFromHttpCookie(c))
	}

	return &cs
}

func (r *response) RawBody() io.ReadCloser {
	if r.raw == nil {
		return nil
	}

	return r.raw.Body
}

func (r *response) Body() []byte {
	if r.raw == nil {
		return nil
	}

	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		panic(err)
	}

	return b
}

func (r *response) String() string {
	return string(r.Body())
}

func (r *response) Size() int64 {
	if r.raw == nil {
		return 0
	}

	return r.raw.ContentLength
}

func (r *response) ReceivedAt() *time.Time {
	return &r.receivedAt
}

func (r *response) DecodeModel() (map[string]interface{}, error) {
	targetModel := make(map[string]interface{})

	if r.raw == nil {
		return nil, fmt.Errorf("no response")
	}

	if r.raw.Body == nil {
		return nil, fmt.Errorf("no body")
	}

	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, fmt.Errorf("empty body")
	}

	err = json.Unmarshal(b, &targetModel)
	if err != nil {
		return nil, err
	}

	return targetModel, nil
}

func (r *response) DecodeCollection() ([]map[string]interface{}, error) {
	targetCollection := make([]map[string]interface{}, 0)

	if r.raw == nil {
		return nil, fmt.Errorf("no response")
	}

	if r.raw.Body == nil {
		return targetCollection, nil
	}

	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return targetCollection, nil
	}

	err = json.Unmarshal(b, &targetCollection)
	if err != nil {
		return nil, err
	}

	return targetCollection, nil
}
