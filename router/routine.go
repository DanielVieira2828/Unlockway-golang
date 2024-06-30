package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetRoutines(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) UpdateRoutine(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) DeleteRoutine(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) CreateRoutine(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
