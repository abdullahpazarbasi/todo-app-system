package infrastructure_adapters_os

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"os"
	"sync"
	drivenAppPortsOs "todo-app-wbff/internal/pkg/app/ports/driven/os"
)

type environmentVariableAccessor struct {
	dotEnvFilePath          string
	dotEnvFileLoaderWrapper sync.Once
}

func NewEnvironmentVariableAccessor(dotEnvFilePath string) drivenAppPortsOs.EnvironmentVariableAccessor {
	return &environmentVariableAccessor{
		dotEnvFilePath: dotEnvFilePath,
	}
}

func (eva *environmentVariableAccessor) NewContextWith(parentContext context.Context) context.Context {
	return context.WithValue(parentContext, drivenAppPortsOs.EnvironmentVariableAccessorKey{}, eva)
}

func (eva *environmentVariableAccessor) Get(key string, defaultValue string) string {
	eva.loadDotEnvIfNotLoaded()
	v, existent := os.LookupEnv(key)
	if existent {
		return v
	}

	return defaultValue
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
