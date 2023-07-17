package driven_app_ports_model

import "context"

type ModelNormalizerKey struct{}

type ModelNormalizer interface {
	NewContextWith(parentContext context.Context) context.Context
	Normalize(sourceModel interface{}) interface{}
	Denormalize(targetModelReference interface{}, sourceModelReference interface{}) error
}
