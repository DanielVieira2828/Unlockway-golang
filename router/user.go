package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) LoginHandler(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}

func (r *Router) RegisterHandler(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
