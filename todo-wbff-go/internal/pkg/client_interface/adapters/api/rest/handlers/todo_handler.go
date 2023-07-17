package client_interface_adapters_rest_api_handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	drivenAppPortsEntity "todo-app-wbff/internal/pkg/app/ports/driven/entity"
	drivingAppPortsTodo "todo-app-wbff/internal/pkg/app/ports/driving/todo"
	clientInterfaceRestApiViews "todo-app-wbff/internal/pkg/client_interface/adapters/api/rest/views"
)

type TodoHandler interface {
	GetCollection(ec echo.Context) error
	Post(ec echo.Context) error
	Patch(ec echo.Context) error
	Delete(ec echo.Context) error
}

type todoHandler struct {
	todoService drivingAppPortsTodo.Service
}

func NewTodoHandler(todoService drivingAppPortsTodo.Service) TodoHandler {
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

	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, err = h.todoService.FindAll(ec.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNoContent, err.Error())
	}

	view := mapTodoCollectionToView(todoCollection)

	return ec.JSON(http.StatusOK, view)
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

	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, err = h.todoService.Add(ec.Request().Context(), userID, value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	view := mapTodoCollectionToView(todoCollection)

	return ec.JSON(http.StatusCreated, view)
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

	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, err = h.todoService.Modify(ec.Request().Context(), userID, id, value, completedRaw)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	view := mapTodoCollectionToView(todoCollection)

	return ec.JSON(http.StatusOK, view)
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

	var todoCollection *[]drivingAppPortsTodo.Todo
	todoCollection, err = h.todoService.Remove(ec.Request().Context(), userID, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	view := mapTodoCollectionToView(todoCollection)

	return ec.JSON(http.StatusOK, view)
}

func resolveUserID(ec echo.Context) (string, error) {
	user := ec.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims.GetSubject()
}

func mapTodoCollectionToView(todoCollection *[]drivingAppPortsTodo.Todo) *clientInterfaceRestApiViews.TodoCollection {
	rs := clientInterfaceRestApiViews.TodoCollection{}
	for _, todo := range *todoCollection {
		rs = append(rs, &clientInterfaceRestApiViews.Todo{
			UserID:    todo.UserID(),
			ID:        todo.(drivenAppPortsEntity.IDProvider).ID(),
			Value:     todo.Label(),
			Completed: todo.IsCompleted(),
		})
	}

	return &rs
}
