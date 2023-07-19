package infrastructure_adapters_restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	drivenAppPortsEncoding "todo-app-wbff/internal/pkg/app/ports/driven/encoding"
)

func resolveBodyReader(body interface{}) io.Reader {
	switch b := body.(type) {
	case nil:
		return nil
	case io.Reader:
		return b
	case string:
		return strings.NewReader(b)
	case *string:
		return strings.NewReader(*b)
	case []byte:
		return bytes.NewReader(b)
	case *[]byte:
		return bytes.NewReader(*b)
	case url.Values:
		return strings.NewReader(b.Encode())
	case *url.Values:
		return strings.NewReader(b.Encode())
	case drivenAppPortsEncoding.Encoder:
		return strings.NewReader(b.Encode())
	case map[string]interface{}, *map[string]interface{}:
		bs, err := json.Marshal(b)
		if err != nil {
			panic(err)
		}

		return bytes.NewReader(bs)
	default:
		panic(fmt.Errorf("invalid request body type %T", b))
	}
}
