package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"todo-app-wbff/configs"
	coreAdapter "todo-app-wbff/internal/pkg/adapter/core"
	drivenAdapterRestful "todo-app-wbff/internal/pkg/adapter/driven/restful"
	drivenAdapterTodo "todo-app-wbff/internal/pkg/adapter/driven/todo"
	drivingAdapterApi "todo-app-wbff/internal/pkg/adapter/driving/api"
	domainFault "todo-app-wbff/internal/pkg/application/domain/fault"
	domainTodo "todo-app-wbff/internal/pkg/application/domain/todo"
	"todo-app-wbff/internal/pkg/application/usecase"
)

func main() {
	ctx := context.Background()

	environmentVariableAccessor := coreAdapter.NewEnvironmentVariableAccessor("configs/.env")

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
	faultFactory := domainFault.NewFactory(
		environmentVariableAccessor,
		configs.EnvironmentVariableNameAppDebug,
	)
	todoFactory := domainTodo.NewFactory()
	drivingAdapterApi.RegisterAPI(
		e,
		environmentVariableAccessor,
		drivingAdapterApi.NewStatusHandler(),
		drivingAdapterApi.NewTokenClaimHandler(
			usecase.NewAuthService(
				faultFactory,
				environmentVariableAccessor,
			),
		),
		drivingAdapterApi.NewTodoHandler(
			usecase.NewTodoService(
				faultFactory,
				todoFactory,
				drivenAdapterTodo.NewRepository(
					faultFactory,
					todoFactory,
					drivenAdapterRestful.NewClient(environmentVariableAccessor.GetOrPanic(configs.EnvironmentVariableNameTodoServiceBaseURL)),
				),
			),
		),
	)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", environmentVariableAccessor.Get(configs.EnvironmentVariableNameHTTPPort, "80"))))
}
