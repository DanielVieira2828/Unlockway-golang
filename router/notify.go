package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetNotify(ctx echo.Context) error {
	notifications := []map[string]interface{}{
		{
			"id":          "1",
			"title":       "Welcome to Unlockway!",
			"description": "Thank you for joining our community. Get started by setting up your profile.",
			"read":        false,
			"date":        "2024-07-09T12:00:00Z",
		},
		{
			"id":          "2",
			"title":       "New Routine Added",
			"description": "A new fitness routine has been added to your account. Check it out now!",
			"read":        true,
			"date":        "2024-07-08T09:30:00Z",
		},
		{
			"id":          "3",
			"title":       "Weekly Summary",
			"description": "Your weekly summary is ready. You have burned 3500 calories this week!",
			"read":        false,
			"date":        "2024-07-07T18:45:00Z",
		},
	}

	return ctx.JSON(http.StatusOK, notifications)
}
