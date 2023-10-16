package driving_adapter_api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	drivingAdapterApiViews "todo-app-service/internal/pkg/adapter/driving/api/views"
	domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"
	domainTodoPort "todo-app-service/internal/pkg/application/domain/todo/port"
	usecasePort "todo-app-service/internal/pkg/application/usecase/port"
)

type TodoHandler interface {
	Post(ec echo.Context) error
	GetCollectionOfUser(ec echo.Context) error
	Put(ec echo.Context) error
	Delete(ec echo.Context) error
}

type todoHandler struct {
	todoService usecasePort.TodoService
}

func NewTodoHandler(todoService usecasePort.TodoService) TodoHandler {
	return &todoHandler{
		todoService: todoService,
	}
}

func (h *todoHandler) Post(ec echo.Context) error {
	var err error

	var form url.Values
	form, err = ec.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var userID string
	userID, err = extractFormParameterValue(form, "user_id", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var label string
	label, err = extractFormParameterValue(form, "label", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var tagKeys []string
	tagKeys, err = extractFormParameterValues(form, "tag", false)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var flt domainFaultPort.Fault
	var id string
	id, flt = h.todoService.Add(ec.Request().Context(), userID, label, tagKeys)
	if flt != nil {
		return flt
	}

	view := drivingAdapterApiViews.IDContainer{
		ID: id,
	}

	return ec.JSON(http.StatusCreated, view)
}

func (h *todoHandler) GetCollectionOfUser(ec echo.Context) error {
	userID := ec.Param("id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user ID must be given")
	}

	var flt domainFaultPort.Fault
	var todoCollection *[]domainTodoPort.TodoEntity
	todoCollection, flt = h.todoService.FindAllForUser(ec.Request().Context(), userID)
	if flt != nil {
		return flt
	}

	view, numberOfTodos := mapTodoEntityCollectionToView(todoCollection)
	if numberOfTodos > 0 {
		return ec.JSON(http.StatusOK, view)
	}

	return ec.NoContent(http.StatusNoContent)
}

func (h *todoHandler) Put(ec echo.Context) error {
	var err error

	id := ec.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID must be given")
	}

	var form url.Values
	form, err = ec.FormParams()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var userID string
	userID, err = extractFormParameterValue(form, "user_id", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var label string
	label, err = extractFormParameterValue(form, "label", true)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var tagKeys []string
	tagKeys, err = extractFormParameterValues(form, "tag", false)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var flt domainFaultPort.Fault
	flt = h.todoService.Modify(ec.Request().Context(), id, userID, label, tagKeys)
	if flt != nil {
		return flt
	}

	return ec.JSON(http.StatusOK, &map[string]interface{}{"ok": true})
}

func (h *todoHandler) Delete(ec echo.Context) error {
	var err error

	id := ec.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID must be given")
	}

	var flt domainFaultPort.Fault
	flt = h.todoService.Remove(ec.Request().Context(), id)
	if flt != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusOK, &map[string]interface{}{"ok": true})
}

func mapTodoEntityCollectionToView(
	collection *[]domainTodoPort.TodoEntity,
) (
	*drivingAdapterApiViews.TodoCollection,
	int,
) {
	var view drivingAdapterApiViews.TodoCollection
	var tagsView drivingAdapterApiViews.TodoTagCollection
	numberOfTodos := 0
	for _, todo := range *collection {
		tagsView = make(drivingAdapterApiViews.TodoTagCollection, 0)
		for _, tag := range *todo.Tags() {
			tagsView = append(tagsView, &drivingAdapterApiViews.TodoTag{
				Key: tag.Key(),
			})
		}
		view = append(view, &drivingAdapterApiViews.Todo{
			UserID: todo.UserID(),
			ID:     todo.ID(),
			Label:  todo.Label(),
			Tags:   &tagsView,
		})
		numberOfTodos++
	}

	return &view, numberOfTodos
}
