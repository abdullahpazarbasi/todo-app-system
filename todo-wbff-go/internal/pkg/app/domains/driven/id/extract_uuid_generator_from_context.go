package driven_app_domains_id

import (
	"context"
	drivenAppPortsId "todo-app-wbff/internal/pkg/app/ports/driven/id"
)

func ExtractUUIDGeneratorFromContext(ctx context.Context) drivenAppPortsId.UUIDGenerator {
	g, existent := ctx.Value(drivenAppPortsId.UUIDGeneratorKey{}).(drivenAppPortsId.UUIDGenerator)
	if existent {
		return g
	}

	panic("UUID generator is not registered in context")
}
