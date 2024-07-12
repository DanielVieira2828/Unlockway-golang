package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetHomeData(ctx echo.Context) error {
	responseData := map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"meals":         2,
			"routines":      2,
			"notifications": 8,
			"weekCalories":  []interface{}{2805.05, 1032.05, 2100.05, 165.05, 4999.05, nil, nil},
		},
	}

	return ctx.JSON(http.StatusOK, responseData)
}
