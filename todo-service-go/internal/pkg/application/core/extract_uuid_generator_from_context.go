package core

import (
	"context"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

func ExtractUUIDGeneratorFromContext(ctx context.Context) corePort.UUIDGenerator {
	g, existent := ctx.Value(corePort.UUIDGeneratorKey{}).(corePort.UUIDGenerator)
	if existent {
		return g
	}

	panic("UUID generator is not registered in context")
}
