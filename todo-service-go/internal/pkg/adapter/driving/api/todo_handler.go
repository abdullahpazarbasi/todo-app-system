package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	drivingAdapterApiViews "todo-app-service/internal/pkg/adapter/driving/api/views"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-service/internal/pkg/application/usecase/port"
)

type TodoHandler interface {
	Post(ec echo.Context) error
	GetCollectionOfUser(ec echo.Context) error
	Patch(ec echo.Context) error
	Delete(ec echo.Context) error
}

type todoHandler struct {
	f           domainFaultPort.Factory
	todoService usecasePort.TodoService
}

func NewTodoHandler(
	faultFactory domainFaultPort.Factory,
	todoService usecasePort.TodoService,
) TodoHandler {
	return &todoHandler{
		f:           faultFactory,
		todoService: todoService,
	}
}

func (h *todoHandler) Post(ec echo.Context) error {
	var err error

	requestBody := drivingAdapterApiViews.Todo{}
	err = ec.Bind(&requestBody)
	if err != nil {
		return h.f.WrapError(
			err,
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
		)
	}

	var flt domainFaultPort.Fault
	var todoID string
	todoID, flt = h.todoService.Add(
		ec.Request().Context(),
		requestBody.UserID,
		requestBody.Label,
		requestBody.TagCollection().Keys(),
	)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.IDContainer{
		ID: todoID,
	}

	return ec.JSON(http.StatusCreated, view)
}

func (h *todoHandler) GetCollectionOfUser(ec echo.Context) error {
	userID := ec.Param("user")
	if userID == "" {
		return h.f.CreateFault(
			h.f.Message("user ID must be given"),
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
		)
	}

	var flt domainFaultPort.Fault
	var todoEntityCollection domainTodoPort.TodoEntityCollection
	todoEntityCollection, flt = h.todoService.FindAllForUser(ec.Request().Context(), userID)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.NewTodoCollectionFromEntityCollection(todoEntityCollection)
	if view.Size() > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Patch(ec echo.Context) error {
	var err error

	todoID := ec.Param("todo")
	if todoID == "" {
		return h.f.CreateFault(
			h.f.Message("todo ID must be given"),
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
		)
	}

	requestBody := drivingAdapterApiViews.Todo{}
	err = ec.Bind(&requestBody)
	if err != nil {
		return h.f.WrapError(
			err,
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
		)
	}

	var flt domainFaultPort.Fault
	flt = h.todoService.Modify(
		ec.Request().Context(),
		todoID,
		requestBody.UserID,
		requestBody.Label,
		requestBody.TagCollection().Keys(),
	)
	if flt != nil {
		return flt
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{"ok": true})
}

func (h *todoHandler) Delete(ec echo.Context) error {
	todoID := ec.Param("todo")
	if todoID == "" {
		return h.f.CreateFault(
			h.f.Message("todo ID must be given"),
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
		)
	}

	var flt domainFaultPort.Fault
	flt = h.todoService.Remove(ec.Request().Context(), todoID)
	if flt != nil {
		return flt
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{"ok": true})
}
