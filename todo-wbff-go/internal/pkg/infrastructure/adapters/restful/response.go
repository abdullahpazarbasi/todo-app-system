package infrastructure_adapters_restful

import (
	"encoding/json"
	"fmt"
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

func (r *response) Header() *map[string][]string {
	if r.raw == nil {
		return nil
	}

	return (*map[string][]string)(&r.raw.Header)
}

func (r *response) Cookies() *[]drivenAppPortsRestful.Cookie {
	cs := make([]drivenAppPortsRestful.Cookie, 0)
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
	var targetModel map[string]interface{}

	if r.raw == nil {
		return targetModel, fmt.Errorf("no response")
	}

	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return targetModel, err
	}

	err = json.Unmarshal(b, &targetModel)
	if err != nil {
		return targetModel, err
	}

	return targetModel, nil
}

func (r *response) DecodeCollection() ([]map[string]interface{}, error) {
	var targetCollection []map[string]interface{}

	if r.raw == nil {
		return targetCollection, fmt.Errorf("no response")
	}

	b, err := io.ReadAll(r.raw.Body)
	if err != nil {
		return targetCollection, err
	}

	err = json.Unmarshal(b, &targetCollection)
	if err != nil {
		return targetCollection, err
	}

	return targetCollection, nil
}
