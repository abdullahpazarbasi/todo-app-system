package driven_app_domains_os

import (
	"context"
	drivenAppPortsOs "todo-app-wbff/internal/pkg/app/ports/driven/os"
)

func ExtractEnvironmentVariableAccessorFromContext(ctx context.Context) drivenAppPortsOs.EnvironmentVariableAccessor {
	eva, existent := ctx.Value(drivenAppPortsOs.EnvironmentVariableAccessorKey{}).(drivenAppPortsOs.EnvironmentVariableAccessor)
	if existent {
		return eva
	}

	panic("environment variable accessor is not registered in context")
}
