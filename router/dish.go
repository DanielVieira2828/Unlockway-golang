package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetDishes(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) UpdateDish(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) CreateDish(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) DeleteDish(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
