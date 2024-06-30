package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetNotify(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
