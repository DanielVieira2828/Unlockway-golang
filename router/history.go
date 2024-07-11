package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetHistory(ctx echo.Context) error {
	histories := []map[string]interface{}{
		{
			"idRoutine":   "routine-uuid-1",
			"routineName": "Healthy Morning Routine",
			"calories":    550.5,
			"date":        "11/07/2024",
			"ingestedMeals": []map[string]interface{}{
				{
					"id":            "history-meal-uuid-1",
					"idMeal":        "meal-uuid-1",
					"idRoutineMeal": "routine-meal-uuid-1",
					"ingested":      true,
					"name":          "Grilled Chicken Salad",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
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
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
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
			"date":        "11/07/2024",
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
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
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
