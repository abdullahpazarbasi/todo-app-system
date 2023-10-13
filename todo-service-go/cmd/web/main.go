package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	coreAdapter "todo-app-service/internal/pkg/adapter/core"
	drivenAdapterDb "todo-app-service/internal/pkg/adapter/driven/db"
	drivingAdapterApi "todo-app-service/internal/pkg/adapter/driving/api"
	domainFault "todo-app-service/internal/pkg/application/domain/fault"
	domainTodo "todo-app-service/internal/pkg/application/domain/todo"
	"todo-app-service/internal/pkg/application/usecase"
)

func main() {
	ctx := context.Background()

	eva := coreAdapter.NewEnvironmentVariableAccessor("configs/.env")
	ctx = eva.NewContextWith(ctx)
	idGenerator := coreAdapter.NewUUIDGenerator()
	ctx = idGenerator.NewContextWith(ctx)

	// Echo instance
	e := echo.New()

	drivingAdapterApi.RegisterMiddlewares(
		e,
		ctx,
	)
	drivingAdapterApi.RegisterStaticAPI(
		e,
		drivingAdapterApi.NewHelloHandler(),
	)

	todoDatabase, err := drivenAdapterDb.ConfigureDB(eva)
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
	faultFactory := domainFault.NewFactory()
	todoService := usecase.NewTodoService(
		todoFactory,
		faultFactory,
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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", eva.Get("HTTP_PORT", "80"))))
}
