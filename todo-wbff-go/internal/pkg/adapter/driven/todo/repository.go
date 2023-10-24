package driven_adapter_todo

import (
	"context"
	"fmt"
	drivenAdapterRestful "todo-app-wbff/internal/pkg/adapter/driven/restful"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-wbff/internal/pkg/application/domain/todo/port"
)

type repository struct {
	f                 domainFaultPort.Factory
	todoFactory       domainTodoPort.Factory
	todoServiceClient drivenAdapterRestful.Client
}

func NewRepository(
	faultFactory domainFaultPort.Factory,
	todoFactory domainTodoPort.Factory,
	todoServiceClient drivenAdapterRestful.Client,
) domainTodoPort.Repository {
	return &repository{
		f:                 faultFactory,
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
		todo.Normalize(),
		drivenAdapterRestful.NewExtraHeaderLineOption("Content-Type", "application/json"),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return "", err
	}

	var rs map[string]interface{}
	rs, err = exchange.Response().DecodeModel()
	if err != nil {
		return "", r.f.WrapError(err)
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
		"/api/users/{user}/todos",
		nil,
		drivenAdapterRestful.NewResourcePathParameterOption("user", userID),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return nil, err
	}

	var rs []map[string]interface{}
	rs, err = exchange.Response().DecodeCollection()
	if err != nil {
		return nil, r.f.WrapError(err)
	}

	todoCollection := r.todoFactory.DenormalizeTodoEntityCollection(&rs)

	return todoCollection, nil
}

func (r *repository) Modify(ctx context.Context, todo domainTodoPort.TodoEntity) error {
	var err error

	id := todo.ID()
	_, err = r.todoServiceClient.Patch(
		ctx,
		"/api/todos/{todo}",
		todo.Normalize(),
		drivenAdapterRestful.NewResourcePathParameterOption("todo", id),
		drivenAdapterRestful.NewExtraHeaderLineOption("Content-Type", "application/json"),
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
		"/api/todos/{todo}",
		drivenAdapterRestful.NewResourcePathParameterOption("todo", id),
		drivenAdapterRestful.NewHTTPErrorHandlingStrategyControllerOption(r.handleHTTPError),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) handleHTTPError(lastExchange drivenAdapterRestful.Exchange, exchangeError error) error {
	if exchangeError != nil {
		return translateHTTPErrorToFault(r.f, exchangeError)
	}
	lastResponse := lastExchange.Response()
	if lastResponse.IsStatusError() {
		responseModel, err := lastResponse.DecodeModel()
		if err != nil {
			return translateHTTPErrorToFault(r.f, err)
		}

		return r.f.DenormalizeError(
			responseModel,
			r.f.ProposedHTTPStatusCode(lastResponse.StatusCode()),
		)
	}

	return nil
}
