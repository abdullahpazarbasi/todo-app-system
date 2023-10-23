package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	drivingAdapterApiViews "todo-app-wbff/internal/pkg/adapter/driving/api/views"
	domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/application/usecase/port"
)

type TodoHandler interface {
	GetCollection(ec echo.Context) error
	Post(ec echo.Context) error
	Patch(ec echo.Context) error
	Delete(ec echo.Context) error
}

type todoHandler struct {
	f           domainFaultPort.Factory
	todoService drivingAppPortsTodo.TodoService
}

func NewTodoHandler(
	faultFactory domainFaultPort.Factory,
	todoService drivingAppPortsTodo.TodoService,
) TodoHandler {
	return &todoHandler{
		f:           faultFactory,
		todoService: todoService,
	}
}

func (h *todoHandler) GetCollection(ec echo.Context) error {
	var err error

	var tokenUserID string
	tokenUserID, err = resolveTokenUserID(ec)
	if err != nil {
		return h.f.WrapError(err, h.f.ProposedHTTPStatusCode(http.StatusUnauthorized))
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.FindAllForUser(ec.Request().Context(), tokenUserID)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.NewTodoCollectionFromModelCollection(todoCollection)
	if view.Size() > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Post(ec echo.Context) error {
	var err error

	var tokenUserID string
	tokenUserID, err = resolveTokenUserID(ec)
	if err != nil {
		return h.f.WrapError(err, h.f.ProposedHTTPStatusCode(http.StatusUnauthorized))
	}

	value := ec.FormValue("value")
	if value == "" {
		return h.f.CreateFault(
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			h.f.Message("value must be given"),
		)
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Add(
		ec.Request().Context(),
		tokenUserID,
		value,
	)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.NewTodoCollectionFromModelCollection(todoCollection)
	if view.Size() > 0 {
		return ec.JSON(http.StatusCreated, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Patch(ec echo.Context) error {
	var err error

	var tokenUserID string
	tokenUserID, err = resolveTokenUserID(ec)
	if err != nil {
		return h.f.WrapError(err, h.f.ProposedHTTPStatusCode(http.StatusUnauthorized))
	}

	todoID := ec.Param("todo")
	if todoID == "" {
		return h.f.CreateFault(
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			h.f.Message("todo ID must be given"),
		)
	}

	completedRaw := ec.FormValue("completed")
	if completedRaw == "" {
		return h.f.CreateFault(
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			h.f.Message("there is no change"),
		)
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Modify(
		ec.Request().Context(),
		tokenUserID,
		todoID,
		completedRaw,
	)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.NewTodoCollectionFromModelCollection(todoCollection)
	if view.Size() > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Delete(ec echo.Context) error {
	var err error

	var tokenUserID string
	tokenUserID, err = resolveTokenUserID(ec)
	if err != nil {
		return h.f.WrapError(err, h.f.ProposedHTTPStatusCode(http.StatusUnauthorized))
	}

	todoID := ec.Param("todo")
	if todoID == "" {
		return h.f.CreateFault(
			h.f.ProposedHTTPStatusCode(http.StatusBadRequest),
			h.f.Message("todo ID must be given"),
		)
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Remove(
		ec.Request().Context(),
		tokenUserID,
		todoID,
	)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.NewTodoCollectionFromModelCollection(todoCollection)
	if view.Size() > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}
