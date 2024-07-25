package handlers

import (
	"fmt"
	"ken/services"
	"ken/views"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func getTodo(id int) (services.Todo, error) {
	for _, todo := range services.Todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return services.Todo{}, fmt.Errorf("no todo with id %d", id)
}

func removeTodo(todos []services.Todo, id int) []services.Todo {
	return append(todos[:id], todos[id+1:]...)
}

func NewTodo(c echo.Context) error {
	name := c.FormValue("todo")
	todo := services.Todo{
		Name: name,
	}
	services.Todos = append(services.Todos, todo)
	return Render(c, http.StatusOK, views.Todo(todo))
}

func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	services.Todos = append(services.Todos[:id], services.Todos[id+1:]...)
	todo, err := getTodo(id)
	if err != nil {
		return err
	}

	return Render(c, http.StatusOK, views.Todo(todo))
}

func ToggleCompleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	services.Todos[id].Complete = !services.Todos[id].Complete
	return Render(c, http.StatusOK, views.Todo(services.Todos[id]))
}
