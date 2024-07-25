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

func getTodo(id int) (services.Todo, int, error) {
	for i, todo := range services.Todos {
		if todo.Id == id {
			return todo, i, nil
		}
	}
	return services.Todo{}, 0, fmt.Errorf("no todo with id %d", id)
}

func removeTodo(todos []services.Todo, id int) ([]services.Todo, error) {
	for i, todo := range services.Todos {
		if todo.Id == id {
			return append(todos[:i], todos[i+1:]...), nil
		}
	}
	return []services.Todo{}, fmt.Errorf("no todo with id %d", id)
}

func NewTodo(c echo.Context) error {
	name := c.FormValue("todo")
	services.Count++
	todo := services.Todo{
		Name: name,
		Id:   services.Count,
	}
	services.Todos = append(services.Todos, todo)
	return Render(c, http.StatusOK, views.Todo(todo))
}

func DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	services.Todos, err = removeTodo(services.Todos, id)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}

func ToggleCompleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	todo, index, err := getTodo(id)
	if err != nil {
		return err
	}
	todo.Complete = !todo.Complete
	services.Todos[index] = todo
	return Render(c, http.StatusOK, views.Todo(todo))
}
