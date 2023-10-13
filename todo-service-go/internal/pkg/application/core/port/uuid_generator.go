package core_port

import "context"

type UUIDGeneratorKey struct{}

type UUIDGenerator interface {
	NewContextWith(parentContext context.Context) context.Context
	GenerateAsString() string
}
