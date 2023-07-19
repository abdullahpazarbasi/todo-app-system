package infrastructure_adapters_id

import (
	"context"
	"github.com/google/uuid"
	drivenAppPortsId "todo-app-wbff/internal/pkg/app/ports/driven/id"
)

type uuidGenerator struct{}

func NewUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) NewContextWith(parentContext context.Context) context.Context {
	return context.WithValue(parentContext, drivenAppPortsId.UUIDGeneratorKey{}, g)
}

func (g *uuidGenerator) GenerateAsString() string {
	return uuid.NewString()
}
