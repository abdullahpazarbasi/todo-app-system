package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	drivenAppDomainsRestful "todo-app-wbff/internal/pkg/app/domains/driven/restful"
	drivenAppDomainsTodo "todo-app-wbff/internal/pkg/app/domains/driven/todo"
	drivingAppDomainsTodo "todo-app-wbff/internal/pkg/app/domains/driving/todo"
	clientInterfaceAdaptersRestApiHandlers "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/handlers"
	clientInterfaceAdaptersRestApiRegistrars "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/registrars"
	infrastructureAdaptersModel "todo-app-wbff/internal/pkg/infrastructure/adapters/model"
	infrastructureAdaptersOs "todo-app-wbff/internal/pkg/infrastructure/adapters/os"
	infrastructureAdaptersRestful "todo-app-wbff/internal/pkg/infrastructure/adapters/restful"
)

func main() {
	ctx := context.Background()

	eva := infrastructureAdaptersOs.NewEnvironmentVariableAccessor("configs/.env")
	ctx = eva.NewContextWith(ctx)
	ctx = infrastructureAdaptersModel.NewModelNormalizer().NewContextWith(ctx)

	// Echo instance
	e := echo.New()

	err := clientInterfaceAdaptersRestApiRegistrars.RegisterMiddlewares(e, ctx)
	if err != nil {
		log.Fatalf("middlewares could not be registered: %v", e)
	}

	helloHandler := clientInterfaceAdaptersRestApiHandlers.NewHelloHandler()
	err = clientInterfaceAdaptersRestApiRegistrars.RegisterStaticAPI(e, helloHandler)
	if err != nil {
		log.Fatalf("static API component could not be registered: %v", e)
	}

	tokenClaimHandler := clientInterfaceAdaptersRestApiHandlers.NewTokenClaimHandler()
	err = clientInterfaceAdaptersRestApiRegistrars.RegisterAuthAPI(e, tokenClaimHandler)
	if err != nil {
		log.Fatalf("auth API component could not be registered: %v", e)
	}

	todoHandler := clientInterfaceAdaptersRestApiHandlers.NewTodoHandler(
		drivingAppDomainsTodo.NewService(
			drivenAppDomainsTodo.NewRepository(
				drivenAppDomainsRestful.NewClientFactoryProvider(
					infrastructureAdaptersRestful.NewClientFactory(),
				),
			),
		),
	)
	err = clientInterfaceAdaptersRestApiRegistrars.RegisterTodoAPI(e, todoHandler, eva)
	if err != nil {
		log.Fatalf("todo API component could not be registered: %v", e)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
