package driven_app_ports_id

import "context"

type UUIDGeneratorKey struct{}

type UUIDGenerator interface {
	NewContextWith(parentContext context.Context) context.Context
	GenerateAsString() string
}
