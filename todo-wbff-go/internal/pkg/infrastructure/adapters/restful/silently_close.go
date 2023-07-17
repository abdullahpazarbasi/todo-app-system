package infrastructure_adapters_restful

import "io"

func silentlyClose(v interface{}) {
	if c, ok := v.(io.Closer); ok {
		silently(c.Close())
	}
}

func silently(_ ...interface{}) {}
