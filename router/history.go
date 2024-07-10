package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetHistory(ctx echo.Context) error {
	histories := []map[string]interface{}{
		{
			"idRoutine":   "routine-uuid-1",
			"routineName": "Healthy Morning Routine",
			"calories":    550.5,
			"date":        time.Now().Format(time.RFC3339Nano),
			"ingestedMeals": []map[string]interface{}{
				{
					"id":            "history-meal-uuid-1",
					"idMeal":        "meal-uuid-1",
					"idRoutineMeal": "routine-meal-uuid-1",
					"ingested":      true,
					"name":          "Grilled Chicken Salad",
					"photo":         "https://example.com/photo.jpg",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
				{
					"id":            "history-meal-uuid-2",
					"idMeal":        "meal-uuid-2",
					"idRoutineMeal": "routine-meal-uuid-2",
					"ingested":      true,
					"name":          "Quinoa Salad",
					"photo":         "https://example.com/photo.jpg",
					"description":   "A nutritious quinoa salad with mixed vegetables.",
					"category":      "Salad",
					"totalCalories": 350.05,
				},
			},
			"days": map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": false, "Saturday": false},
		},
		{
			"idRoutine":   "routine-uuid-2",
			"routineName": "Evening Fitness Routine",
			"calories":    600.05,
			"date":        time.Now().Format(time.RFC3339Nano),
			"ingestedMeals": []map[string]interface{}{
				{
					"id":            "history-meal-uuid-3",
					"idMeal":        "meal-uuid-3",
					"idRoutineMeal": "routine-meal-uuid-3",
					"ingested":      true,
					"name":          "Broccoli and Beef Stir-fry",
					"photo":         "https://example.com/photo.jpg",
					"description":   "A delicious stir-fry with beef and broccoli.",
					"category":      "Stir-fry",
					"totalCalories": 400.05,
				},
				{
					"id":            "history-meal-uuid-4",
					"idMeal":        "meal-uuid-1",
					"idRoutineMeal": "routine-meal-uuid-4",
					"ingested":      true,
					"name":          "Grilled Chicken Salad",
					"photo":         "https://example.com/photo.jpg",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
			},
			"days": map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": true, "Saturday": false},
		},
	}

	return ctx.JSON(http.StatusOK, histories)
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
