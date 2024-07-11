package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type RoutineMealOnGet struct {
	ID            string  `json:"id"`
	NotifyAt      string  `json:"notifyAt"`
	MealID        string  `json:"mealId"`
	Photo         string  `json:"photo"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	TotalCalories float64 `json:"totalCalories"`
}

type RoutineModel struct {
	ID                    string             `json:"id"`
	Name                  string             `json:"name"`
	InUsage               bool               `json:"inUsage"`
	Meals                 []RoutineMealOnGet `json:"meals"`
	WeekRepetitions       []bool             `json:"weekRepetitions"`
	TotalCaloriesInTheDay float64            `json:"totalCaloriesInTheDay"`
	CreatedAt             time.Time          `json:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt"`
}

func (r *Router) GetRoutines(ctx echo.Context) error {
	routines := []map[string]interface{}{
		{
			"id":      "routine-uuid-1",
			"name":    "Healthy Morning Routine",
			"inUsage": true,
			"meals": []map[string]interface{}{
				{
					"id":            "routine-meal-uuid-1",
					"notifyAt":      "07:30",
					"mealId":        "meal-uuid-1",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Grilled Chicken Salad",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
				{
					"id":            "routine-meal-uuid-2",
					"notifyAt":      "12:30",
					"mealId":        "meal-uuid-2",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Quinoa Salad",
					"description":   "A nutritious quinoa salad with mixed vegetables.",
					"category":      "Salad",
					"totalCalories": 350.05,
				},
			},
			"weekRepetitions":       map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": false, "Saturday": false},
			"totalCaloriesInTheDay": 550.05,
			"createdAt":             time.Now().Format(time.RFC3339Nano),
			"updatedAt":             time.Now().Format(time.RFC3339Nano),
		},
		{
			"id":      "routine-uuid-2",
			"name":    "Evening Fitness Routine",
			"inUsage": true,
			"meals": []map[string]interface{}{
				{
					"id":            "routine-meal-uuid-3",
					"notifyAt":      "13:00",
					"mealId":        "meal-uuid-1",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Grilled Chicken Salad",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
				{
					"id":            "routine-meal-uuid-4",
					"notifyAt":      "19:30",
					"mealId":        "meal-uuid-3",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Broccoli and Beef Stir-fry",
					"description":   "A delicious stir-fry with beef and broccoli.",
					"category":      "Stir-fry",
					"totalCalories": 400.05,
				},
			},
			"weekRepetitions":       map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": true, "Saturday": false},
			"totalCaloriesInTheDay": 600.05,
			"createdAt":             time.Now().Format(time.RFC3339Nano),
			"updatedAt":             time.Now().Format(time.RFC3339Nano),
		},
	}

	return ctx.JSON(http.StatusOK, routines)
}

func (r *Router) GetOnUseRoutines(ctx echo.Context) error {
	routines := []map[string]interface{}{
		{
			"id":      "routine-uuid-1",
			"name":    "Healthy Morning Routine",
			"inUsage": true,
			"meals": []map[string]interface{}{
				{
					"id":            "routine-meal-uuid-1",
					"notifyAt":      "07:30",
					"mealId":        "meal-uuid-1",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Grilled Chicken Salad",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
				{
					"id":            "routine-meal-uuid-2",
					"notifyAt":      "12:30",
					"mealId":        "meal-uuid-2",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Quinoa Salad",
					"description":   "A nutritious quinoa salad with mixed vegetables.",
					"category":      "Salad",
					"totalCalories": 350.05,
				},
			},
			"weekRepetitions":       map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": false, "Saturday": false},
			"totalCaloriesInTheDay": 550.05,
			"createdAt":             time.Now().Format(time.RFC3339Nano),
			"updatedAt":             time.Now().Format(time.RFC3339Nano),
		},
		{
			"id":      "routine-uuid-2",
			"name":    "Evening Fitness Routine",
			"inUsage": true,
			"meals": []map[string]interface{}{
				{
					"id":            "routine-meal-uuid-3",
					"notifyAt":      "13:00",
					"mealId":        "meal-uuid-1",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Grilled Chicken Salad",
					"description":   "A healthy grilled chicken salad with fresh vegetables.",
					"category":      "Salad",
					"totalCalories": 200.05,
				},
				{
					"id":            "routine-meal-uuid-4",
					"notifyAt":      "19:30",
					"mealId":        "meal-uuid-3",
					"photo":         "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
					"name":          "Broccoli and Beef Stir-fry",
					"description":   "A delicious stir-fry with beef and broccoli.",
					"category":      "Stir-fry",
					"totalCalories": 400.05,
				},
			},
			"weekRepetitions":       map[string]bool{"Sunday": true, "Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true, "Friday": true, "Saturday": false},
			"totalCaloriesInTheDay": 600.05,
			"createdAt":             time.Now().Format(time.RFC3339Nano),
			"updatedAt":             time.Now().Format(time.RFC3339Nano),
		},
	}

	return ctx.JSON(http.StatusOK, routines)
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
