package driven_app_domains_model

import (
	"context"
	drivenAppPortsModel "todo-app-wbff/internal/pkg/app/ports/driven/model"
)

func ExtractModelNormalizerFromContext(ctx context.Context) drivenAppPortsModel.ModelNormalizer {
	n, existent := ctx.Value(drivenAppPortsModel.ModelNormalizerKey{}).(drivenAppPortsModel.ModelNormalizer)
	if existent {
		return n
	}

	panic("model normalizer is not registered in context")
}
