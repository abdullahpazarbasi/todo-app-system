package core_adapter

import (
	"context"
	"github.com/google/uuid"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

type uuidGenerator struct{}

func NewUUIDGenerator() corePort.UUIDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) NewContextWith(parentContext context.Context) context.Context {
	return context.WithValue(parentContext, corePort.UUIDGeneratorKey{}, g)
}

func (g *uuidGenerator) GenerateAsString() string {
	return uuid.NewString()
}
