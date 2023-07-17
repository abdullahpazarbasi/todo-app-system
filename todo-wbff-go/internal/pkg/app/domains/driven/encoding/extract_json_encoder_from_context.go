package driven_app_domains_encoding

import (
	"context"
	drivenAppPortsEncoding "todo-app-wbff/internal/pkg/app/ports/driven/encoding"
)

func ExtractJSONEncoderFromContext(ctx context.Context) drivenAppPortsEncoding.Encoder {
	e, existent := ctx.Value(drivenAppPortsEncoding.JSONEncoderKey{}).(drivenAppPortsEncoding.Encoder)
	if existent {
		return e
	}

	panic("JSON encoder is not registered in context")
}
