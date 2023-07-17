package infrastructure_adapters_restful

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

type response struct {
	raw        *http.Response
	receivedAt time.Time
}

func (r *response) Proto() string {
	return r.raw.Proto
}

func (r *response) StatusCode() int {
	return r.raw.StatusCode
}

func (r *response) Status() string {
	return r.raw.Status
}

func (r *response) IsStatusSuccess() bool {
	return r.raw.StatusCode > 199 && r.raw.StatusCode < 300
}

func (r *response) IsStatusFail() bool {
	return r.raw.StatusCode > 399
}

func (r *response) Header() *map[string][]string {
	return (*map[string][]string)(&r.raw.Header)
}

func (r *response) Cookies() *[]drivenAppPortsRestful.Cookie {
	cs := make([]drivenAppPortsRestful.Cookie, 0)
	for _, c := range r.raw.Cookies() {
		cs = append(cs, NewCookieFromHttpCookie(c))
	}

	return &cs
}

func (r *response) RawBody() io.ReadCloser {
	return r.raw.Body
}

func (r *response) Body() []byte {
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
	return r.raw.ContentLength
}

func (r *response) ReceivedAt() *time.Time {
	return &r.receivedAt
}

func (r *response) DecodeModel() (map[string]interface{}, error) {
	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	var targetModel map[string]interface{}
	err = json.Unmarshal(b, &targetModel)
	if err != nil {
		return nil, err
	}

	return targetModel, nil
}

func (r *response) DecodeCollection() ([]map[string]interface{}, error) {
	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	var targetCollection []map[string]interface{}
	err = json.Unmarshal(b, &targetCollection)
	if err != nil {
		return nil, err
	}

	return targetCollection, nil
}
