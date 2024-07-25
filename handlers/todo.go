package handlers

import (
	"ken/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func NewTodo(c echo.Context) error {
	todo := c.FormValue("todo")
	return Render(c, http.StatusOK, views.Todo(todo, 1, true))
}
