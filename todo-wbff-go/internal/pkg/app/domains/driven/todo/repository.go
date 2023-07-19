package driven_app_domains_todo

import (
	"context"
	"fmt"
	drivenAppDomainsModel "todo-app-wbff/internal/pkg/app/domains/driven/model"
	drivenAppDomainsOs "todo-app-wbff/internal/pkg/app/domains/driven/os"
	drivenAppDomainsRestful "todo-app-wbff/internal/pkg/app/domains/driven/restful"
	drivenAppPortsCore "todo-app-wbff/internal/pkg/app/ports/driven/core"
	drivenAppPortsEntity "todo-app-wbff/internal/pkg/app/ports/driven/entity"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/app/ports/driving/todo"
)

type Repository interface {
	Add(ctx context.Context, todo drivingAppPortsTodo.Todo) (string, error)
	FindAllForUser(ctx context.Context, userID string) (*[]*TodoEntity, error)
	Replace(ctx context.Context, todo drivingAppPortsTodo.Todo) error
	Remove(ctx context.Context, id string) error
}

type repository struct {
	restfulClientFactoryProvider drivenAppPortsRestful.ClientFactoryProvider
}

func NewRepository(restfulClientFactoryProvider drivenAppPortsRestful.ClientFactoryProvider) *repository {
	return &repository{
		restfulClientFactoryProvider: restfulClientFactoryProvider,
	}
}

func (r *repository) Add(ctx context.Context, todo drivingAppPortsTodo.Todo) (string, error) {
	mn := drivenAppDomainsModel.ExtractModelNormalizerFromContext(ctx)
	eva := drivenAppDomainsOs.ExtractEnvironmentVariableAccessorFromContext(ctx)
	c := r.restfulClientFactoryProvider.Provide().Create(
		eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTodoServiceBaseURL),
	)
	exchange, err := c.Post(ctx, "/api/todos", mn.Normalize(todo))
	if err != nil {
		return "", err
	}

	var rs map[string]interface{}
	rs, err = exchange.Response().DecodeModel()
	if err != nil {
		return "", err
	}

	id, existent := rs["id"]
	if !existent {
		return "", fmt.Errorf("malformed response, id is not responded")
	}

	return id.(string), nil
}

func (r *repository) FindAllForUser(ctx context.Context, userID string) (*[]*TodoEntity, error) {
	mn := drivenAppDomainsModel.ExtractModelNormalizerFromContext(ctx)
	eva := drivenAppDomainsOs.ExtractEnvironmentVariableAccessorFromContext(ctx)
	c := r.restfulClientFactoryProvider.Provide().Create(
		eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTodoServiceBaseURL),
	)
	exchange, err := c.Get(
		ctx,
		"/api/users/{id}/todos",
		nil,
		drivenAppDomainsRestful.NewResourcePathParameterOption("id", userID),
	)
	if err != nil {
		return nil, err
	}

	var rs []map[string]interface{}
	rs, err = exchange.Response().DecodeCollection()
	if err != nil {
		return nil, err
	}

	var todoCollection []*TodoEntity
	err = mn.Denormalize(&todoCollection, rs)
	if err != nil {
		return nil, err
	}

	return &todoCollection, nil
}

func (r *repository) Replace(ctx context.Context, todo drivingAppPortsTodo.Todo) error {
	id := todo.(drivenAppPortsEntity.IDProvider).ID()
	mn := drivenAppDomainsModel.ExtractModelNormalizerFromContext(ctx)
	eva := drivenAppDomainsOs.ExtractEnvironmentVariableAccessorFromContext(ctx)
	c := r.restfulClientFactoryProvider.Provide().Create(
		eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTodoServiceBaseURL),
	)
	_, err := c.Put(
		ctx,
		"/api/todos/{id}",
		mn.Normalize(todo),
		drivenAppDomainsRestful.NewResourcePathParameterOption("id", id),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Remove(ctx context.Context, id string) error {
	eva := drivenAppDomainsOs.ExtractEnvironmentVariableAccessorFromContext(ctx)
	c := r.restfulClientFactoryProvider.Provide().Create(
		eva.GetOrPanic(drivenAppPortsCore.EnvironmentVariableNameTodoServiceBaseURL),
	)
	_, err := c.Delete(
		ctx,
		"/api/todos/{id}",
		drivenAppDomainsRestful.NewResourcePathParameterOption("id", id),
	)
	if err != nil {
		return err
	}

	return nil
}
