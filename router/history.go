package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetHistory(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) UpdateHistory(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) DeleteHistory(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) CreateHistory(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
