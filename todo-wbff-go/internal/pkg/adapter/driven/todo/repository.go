package driven_adapter_todo

import (
	"context"
	"fmt"
	drivenAdapterRestful "todo-app-wbff/internal/pkg/adapter/driven/restful"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
)

type repository struct {
	faultFactory      domainFaultPort.Factory
	todoFactory       domainTodoPort.Factory
	todoServiceClient drivenAdapterRestful.Client
}

func NewRepository(
	faultFactory domainFaultPort.Factory,
	todoFactory domainTodoPort.Factory,
	todoServiceClient drivenAdapterRestful.Client,
) domainTodoPort.Repository {
	return &repository{
		faultFactory:      faultFactory,
		todoFactory:       todoFactory,
		todoServiceClient: todoServiceClient,
	}
}

func (r *repository) Add(ctx context.Context, todo domainTodoPort.TodoEntity) (string, error) {
	var err error

	var exchange drivenAdapterRestful.Exchange
	exchange, err = r.todoServiceClient.Post(
		ctx,
		"/api/todos",
		normalizeTodoEntity(todo),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
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

func (r *repository) FindAllForUser(ctx context.Context, userID string) (*[]domainTodoPort.TodoEntity, error) {
	exchange, err := r.todoServiceClient.Get(
		ctx,
		"/api/users/{id}/todos",
		nil,
		drivenAdapterRestful.NewResourcePathParameterOption("id", userID),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return nil, err
	}

	var rs []map[string]interface{}
	rs, err = exchange.Response().DecodeCollection()
	if err != nil {
		return nil, err
	}

	var todoCollection []domainTodoPort.TodoEntity
	err = denormalizeToTodoEntityCollection(&todoCollection, &rs, r.todoFactory)
	if err != nil {
		return nil, err
	}

	return &todoCollection, nil
}

func (r *repository) Replace(ctx context.Context, todo domainTodoPort.TodoEntity) error {
	var err error

	id := todo.ID()
	_, err = r.todoServiceClient.Put(
		ctx,
		"/api/todos/{id}",
		normalizeTodoEntity(todo),
		drivenAdapterRestful.NewResourcePathParameterOption("id", id),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Remove(ctx context.Context, id string) error {
	_, err := r.todoServiceClient.Delete(
		ctx,
		"/api/todos/{id}",
		drivenAdapterRestful.NewResourcePathParameterOption("id", id),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) handleHTTPError(lastExchange drivenAdapterRestful.Exchange, cause error) error {
	if cause != nil {
		return nil
	}
	lastResponse := lastExchange.Response()
	if lastResponse.IsStatusError() {
		responseModel, err := lastResponse.DecodeModel()
		if err != nil {
			return err
		}

		return r.faultFactory.DenormalizeError(
			&responseModel,
			r.faultFactory.ProposedHTTPStatusCode(lastResponse.StatusCode()),
		)
	}

	return nil
}
