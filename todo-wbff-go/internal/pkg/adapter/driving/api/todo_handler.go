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
	todoService drivingAppPortsTodo.TodoService
}

func NewTodoHandler(todoService drivingAppPortsTodo.TodoService) TodoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func (h *todoHandler) GetCollection(ec echo.Context) error {
	var err error

	var userID string
	userID, err = resolveUserID(ec)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.FindAllForUser(ec.Request().Context(), userID)
	if flt != nil {
		return flt
	}

	view, sizeOfCollection := mapTodoEntityCollectionToView(todoCollection)
	if sizeOfCollection > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Post(ec echo.Context) error {
	var err error

	var userID string
	userID, err = resolveUserID(ec)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	value := ec.FormValue("value")
	if value == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "value must be given")
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Add(ec.Request().Context(), userID, value)
	if flt != nil {
		return flt
	}

	view, sizeOfCollection := mapTodoEntityCollectionToView(todoCollection)
	if sizeOfCollection > 0 {
		return ec.JSON(http.StatusCreated, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Patch(ec echo.Context) error {
	var err error

	var userID string
	userID, err = resolveUserID(ec)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	id := ec.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be given")
	}

	value := ec.FormValue("value")
	completedRaw := ec.FormValue("completed")
	if value == "" && completedRaw == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "there is no change")
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Modify(ec.Request().Context(), userID, id, value, completedRaw)
	if flt != nil {
		return flt
	}

	view, sizeOfCollection := mapTodoEntityCollectionToView(todoCollection)
	if sizeOfCollection > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Delete(ec echo.Context) error {
	var err error

	var userID string
	userID, err = resolveUserID(ec)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	id := ec.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be given")
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, flt = h.todoService.Remove(ec.Request().Context(), userID, id)
	if flt != nil {
		return flt
	}

	view, sizeOfCollection := mapTodoEntityCollectionToView(todoCollection)
	if sizeOfCollection > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func mapTodoEntityCollectionToView(
	todoCollection *[]drivingAppPortsTodo.Todo,
) (
	*drivingAdapterApiViews.TodoCollection,
	int,
) {
	sizeOfCollection := 0
	rs := drivingAdapterApiViews.TodoCollection{}
	for _, todo := range *todoCollection {
		rs = append(rs, &drivingAdapterApiViews.Todo{
			UserID:    todo.UserID(),
			ID:        todo.ID(),
			Value:     todo.Value(),
			Completed: todo.IsCompleted(),
		})
		sizeOfCollection++
	}

	return &rs, sizeOfCollection
}
