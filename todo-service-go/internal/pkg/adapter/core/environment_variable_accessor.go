package core_adapter

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
	"sync"
	corePort "todo-app-service/internal/pkg/application/core/port"
)

type environmentVariableAccessor struct {
	dotEnvFilePath          string
	dotEnvFileLoaderWrapper sync.Once
}

func NewEnvironmentVariableAccessor(dotEnvFilePath string) corePort.EnvironmentVariableAccessor {
	return &environmentVariableAccessor{
		dotEnvFilePath: dotEnvFilePath,
	}
}

func (eva *environmentVariableAccessor) Get(key string, defaultValue string) string {
	eva.loadDotEnvIfNotLoaded()
	v, existent := os.LookupEnv(key)
	if existent {
		return v
	}

	return defaultValue
}

func (eva *environmentVariableAccessor) GetOrThrowError(key string) (string, error) {
	eva.loadDotEnvIfNotLoaded()
	v, existent := os.LookupEnv(key)
	if existent {
		return v, nil
	}

	return "", fmt.Errorf("environment variable '%s' does not exist", key)
}

func (eva *environmentVariableAccessor) GetOrPanic(key string) string {
	eva.loadDotEnvIfNotLoaded()
	v, existent := os.LookupEnv(key)
	if existent {
		return v
	}

	panic(fmt.Sprintf("environment variable '%s' does not exist", key))
}

func (eva *environmentVariableAccessor) loadDotEnvIfNotLoaded() {
	eva.dotEnvFileLoaderWrapper.Do(func() {
		if eva.dotEnvFilePath == "" {
			return
		}
		err := godotenv.Load(eva.dotEnvFilePath)
		if err != nil {
			log.Fatalf(".env file could not be loaded: %v", err)
		}
	})
}
