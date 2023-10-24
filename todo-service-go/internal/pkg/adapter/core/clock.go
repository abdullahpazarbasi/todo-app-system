package core_adapter

import (
	"time"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

type clock struct {
}

func NewClock() corePort.Clock {
	return &clock{}
}

func (c *clock) Now() *time.Time {
	now := time.Now()

	return &now
}
