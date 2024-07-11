package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetDishes(ctx echo.Context) error {
	meals := []map[string]interface{}{
		{
			"id":                "meal-uuid-1",
			"name":              "Grilled Chicken Salad",
			"photo":             "https://assets.epicurious.com/photos/64a845e67799ee8651e4fb8f/16:9/w_6815,h_3833,c_limit/AshaGrilledChickenSalad_RECIPE_070523_56498.jpg",
			"description":       "A healthy grilled chicken salad with fresh vegetables.",
			"category":          "LUNCH",
			"preparationMethod": "Grill the chicken and mix with fresh vegetables.",
			"ingredients": []map[string]interface{}{
				{"id": "ingredient-uuid-1", "name": "Chicken Breast", "measure": "GRAMS", "amount": 200.05},
				{"id": "ingredient-uuid-2", "name": "Lettuce", "measure": "GRAMS", "amount": 50.0},
				{"id": "ingredient-uuid-3", "name": "Tomato", "measure": "GRAMS", "amount": 30.0},
				{"id": "ingredient-uuid-4", "name": "Cucumber", "measure": "GRAMS", "amount": 30.0},
				{"id": "ingredient-uuid-5", "name": "Olive Oil", "measure": "MILILITERS", "amount": 10.0},
			},
			"totalCalories": 350.0,
			"createdAt":     time.Now().Format(time.RFC3339Nano),
			"updatedAt":     time.Now().Format(time.RFC3339Nano),
		},
		{
			"id":                "meal-uuid-2",
			"name":              "Vegetable Stir Fry",
			"photo":             "https://www.wholesomeyum.com/wp-content/uploads/2020/11/wholesomeyum-Stir-Fry-Vegetables-15.jpg",
			"description":       "A delicious vegetable stir fry with a mix of fresh vegetables.",
			"category":          "DINNER",
			"preparationMethod": "Stir fry the vegetables with soy sauce.",
			"ingredients": []map[string]interface{}{
				{"id": "ingredient-uuid-6", "name": "Bell Pepper", "measure": "GRAMS", "amount": 100.0},
				{"id": "ingredient-uuid-7", "name": "Broccoli", "measure": "GRAMS", "amount": 100.0},
				{"id": "ingredient-uuid-8", "name": "Carrot", "measure": "GRAMS", "amount": 50.0},
				{"id": "ingredient-uuid-9", "name": "Soy Sauce", "measure": "MILILITERS", "amount": 20.0},
			},
			"totalCalories": 150.0,
			"createdAt":     time.Now().Format(time.RFC3339Nano),
			"updatedAt":     time.Now().Format(time.RFC3339Nano),
		},
	}

	return ctx.JSON(http.StatusOK, meals)
}

func (r *Router) GetIngredients(ctx echo.Context) error {
	meals := []map[string]interface{}{
		{
			"id":          "ingredient-uuid-1",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Chicken Breast",
			"measure":     "GRAMS",
			"description": "High-protein chicken breast.",
			"calories":    165.05,
			"proteins":    31.05,
			"water":       65.05,
			"minerals":    "Iron, Zinc",
			"vitamins":    "B6, B12",
			"fats":        3.6,
		},
		{
			"id":          "ingredient-uuid-2",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Lettuce",
			"measure":     "GRAMS",
			"description": "Fresh and crispy lettuce.",
			"calories":    15.05,
			"proteins":    1.4,
			"water":       95.05,
			"minerals":    "Calcium, Magnesium",
			"vitamins":    "A, C, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-3",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Tomato",
			"measure":     "GRAMS",
			"description": "Juicy and ripe tomatoes.",
			"calories":    18.05,
			"proteins":    0.9,
			"water":       94.05,
			"minerals":    "Potassium",
			"vitamins":    "C, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-4",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Cucumber",
			"measure":     "GRAMS",
			"description": "Refreshing cucumber.",
			"calories":    16.05,
			"proteins":    0.7,
			"water":       96.05,
			"minerals":    "Calcium, Magnesium",
			"vitamins":    "C, K",
			"fats":        0.1,
		},
		{
			"id":          "ingredient-uuid-5",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Olive Oil",
			"measure":     "MILILITERS",
			"description": "Healthy olive oil.",
			"calories":    120.05,
			"proteins":    0.05,
			"water":       0.05,
			"minerals":    "Iron",
			"vitamins":    "E",
			"fats":        14.05,
		},
		{
			"id":          "ingredient-uuid-6",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Bell Pepper",
			"measure":     "GRAMS",
			"description": "Colorful bell pepper.",
			"calories":    26.05,
			"proteins":    1.05,
			"water":       92.05,
			"minerals":    "Calcium, Iron",
			"vitamins":    "A, C",
			"fats":        0.3,
		},
		{
			"id":          "ingredient-uuid-7",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Broccoli",
			"measure":     "GRAMS",
			"description": "Fresh broccoli.",
			"calories":    34.05,
			"proteins":    2.8,
			"water":       90.05,
			"minerals":    "Iron, Magnesium",
			"vitamins":    "A, C, K",
			"fats":        0.4,
		},
		{
			"id":          "ingredient-uuid-8",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Carrot",
			"measure":     "GRAMS",
			"description": "Crunchy carrot.",
			"calories":    41.05,
			"proteins":    0.9,
			"water":       88.05,
			"minerals":    "Potassium",
			"vitamins":    "A, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-9",
			"photo":       "https://i0.wp.com/mercadoeconsumo.com.br/wp-content/uploads/2019/04/Que-comida-saud%C3%A1vel-que-nada-brasileiro-gosta-de-fast-food.jpg?fit=1600%2C1067&ssl=1",
			"name":        "Soy Sauce",
			"measure":     "MILILITERS",
			"description": "Savory soy sauce.",
			"calories":    53.05,
			"proteins":    8.05,
			"water":       67.05,
			"minerals":    "Sodium",
			"vitamins":    "B6",
			"fats":        0.1,
		},
	}

	return ctx.JSON(http.StatusOK, meals)
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
