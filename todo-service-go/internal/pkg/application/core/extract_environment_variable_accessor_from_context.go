package core

import (
	"context"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

func ExtractEnvironmentVariableAccessorFromContext(ctx context.Context) corePort.EnvironmentVariableAccessor {
	eva, existent := ctx.Value(corePort.EnvironmentVariableAccessorKey{}).(corePort.EnvironmentVariableAccessor)
	if existent {
		return eva
	}

	panic("environment variable accessor is not registered in context")
}
