package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"todo-app-service/configs"
	coreAdapter "todo-app-service/internal/pkg/adapter/core"
	drivenAdapterDb "todo-app-service/internal/pkg/adapter/driven/db"
	drivingAdapterApi "todo-app-service/internal/pkg/adapter/driving/api"
	domainFault "todo-app-service/internal/pkg/application/domain/fault"
	domainTodo "todo-app-service/internal/pkg/application/domain/todo"
	"todo-app-service/internal/pkg/application/usecase"
)

func main() {
	ctx := context.Background()

	environmentVariableAccessor := coreAdapter.NewEnvironmentVariableAccessor("configs/.env")
	idGenerator := coreAdapter.NewUUIDGenerator()

	// Echo instance
	e := echo.New()

	drivingAdapterApi.RegisterMiddlewares(
		e,
		environmentVariableAccessor,
		ctx,
	)
	drivingAdapterApi.RegisterStaticAPI(
		e,
		drivingAdapterApi.NewHelloHandler(),
	)

	todoDatabase, err := drivenAdapterDb.ConfigureDB(environmentVariableAccessor)
	if err != nil {
		e.Logger.Fatal(err)
	}
	var todoDatabaseClient *sql.DB
	todoDatabaseClient, err = todoDatabase.DB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		_ = todoDatabaseClient.Close()
	}()

	todoFactory := domainTodo.NewFactory(idGenerator)
	faultFactory := domainFault.NewFactory(
		environmentVariableAccessor,
		configs.EnvironmentVariableNameAppDebug,
	)
	todoService := usecase.NewTodoService(
		todoFactory,
		faultFactory,
		idGenerator,
		drivenAdapterDb.NewRepository(
			todoDatabase,
			todoFactory,
			faultFactory,
		),
	)
	drivingAdapterApi.RegisterAPIs(
		e,
		drivingAdapterApi.NewStatusHandler(
			todoDatabaseClient,
		),
		drivingAdapterApi.NewTodoHandler(
			todoService,
		),
	)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", environmentVariableAccessor.Get(configs.EnvironmentVariableNameHTTPPort, "80"))))
}
