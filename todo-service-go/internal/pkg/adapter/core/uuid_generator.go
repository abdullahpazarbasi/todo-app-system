package core_adapter

import (
	"github.com/google/uuid"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

type uuidGenerator struct{}

func NewUUIDGenerator() corePort.UUIDGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) GenerateAsString() string {
	return uuid.NewString()
}
