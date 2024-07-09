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
			"category":          "Almoço",
			"preparationMethod": "Grill the chicken and mix with fresh vegetables.",
			"ingredients": []map[string]interface{}{
				{"id": "ingredient-uuid-1", "name": "Chicken Breast", "measure": "grams", "amount": 200.0},
				{"id": "ingredient-uuid-2", "name": "Lettuce", "measure": "grams", "amount": 50.0},
				{"id": "ingredient-uuid-3", "name": "Tomato", "measure": "grams", "amount": 30.0},
				{"id": "ingredient-uuid-4", "name": "Cucumber", "measure": "grams", "amount": 30.0},
				{"id": "ingredient-uuid-5", "name": "Olive Oil", "measure": "ml", "amount": 10.0},
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
			"category":          "Jantar",
			"preparationMethod": "Stir fry the vegetables with soy sauce.",
			"ingredients": []map[string]interface{}{
				{"id": "ingredient-uuid-6", "name": "Bell Pepper", "measure": "grams", "amount": 100.0},
				{"id": "ingredient-uuid-7", "name": "Broccoli", "measure": "grams", "amount": 100.0},
				{"id": "ingredient-uuid-8", "name": "Carrot", "measure": "grams", "amount": 50.0},
				{"id": "ingredient-uuid-9", "name": "Soy Sauce", "measure": "ml", "amount": 20.0},
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
			"photo":       "https://example.com/chicken_breast.jpg",
			"name":        "Chicken Breast",
			"measure":     "grams",
			"description": "High-protein chicken breast.",
			"calories":    165.0,
			"proteins":    31.0,
			"water":       65.0,
			"minerals":    "Iron, Zinc",
			"vitamins":    "B6, B12",
			"fats":        3.6,
		},
		{
			"id":          "ingredient-uuid-2",
			"photo":       "https://example.com/lettuce.jpg",
			"name":        "Lettuce",
			"measure":     "grams",
			"description": "Fresh and crispy lettuce.",
			"calories":    15.0,
			"proteins":    1.4,
			"water":       95.0,
			"minerals":    "Calcium, Magnesium",
			"vitamins":    "A, C, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-3",
			"photo":       "https://example.com/tomato.jpg",
			"name":        "Tomato",
			"measure":     "grams",
			"description": "Juicy and ripe tomatoes.",
			"calories":    18.0,
			"proteins":    0.9,
			"water":       94.0,
			"minerals":    "Potassium",
			"vitamins":    "C, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-4",
			"photo":       "https://example.com/cucumber.jpg",
			"name":        "Cucumber",
			"measure":     "grams",
			"description": "Refreshing cucumber.",
			"calories":    16.0,
			"proteins":    0.7,
			"water":       96.0,
			"minerals":    "Calcium, Magnesium",
			"vitamins":    "C, K",
			"fats":        0.1,
		},
		{
			"id":          "ingredient-uuid-5",
			"photo":       "https://example.com/olive_oil.jpg",
			"name":        "Olive Oil",
			"measure":     "ml",
			"description": "Healthy olive oil.",
			"calories":    120.0,
			"proteins":    0.0,
			"water":       0.0,
			"minerals":    "Iron",
			"vitamins":    "E",
			"fats":        14.0,
		},
		{
			"id":          "ingredient-uuid-6",
			"photo":       "https://example.com/bell_pepper.jpg",
			"name":        "Bell Pepper",
			"measure":     "grams",
			"description": "Colorful bell pepper.",
			"calories":    26.0,
			"proteins":    1.0,
			"water":       92.0,
			"minerals":    "Calcium, Iron",
			"vitamins":    "A, C",
			"fats":        0.3,
		},
		{
			"id":          "ingredient-uuid-7",
			"photo":       "https://example.com/broccoli.jpg",
			"name":        "Broccoli",
			"measure":     "grams",
			"description": "Fresh broccoli.",
			"calories":    34.0,
			"proteins":    2.8,
			"water":       90.0,
			"minerals":    "Iron, Magnesium",
			"vitamins":    "A, C, K",
			"fats":        0.4,
		},
		{
			"id":          "ingredient-uuid-8",
			"photo":       "https://example.com/carrot.jpg",
			"name":        "Carrot",
			"measure":     "grams",
			"description": "Crunchy carrot.",
			"calories":    41.0,
			"proteins":    0.9,
			"water":       88.0,
			"minerals":    "Potassium",
			"vitamins":    "A, K",
			"fats":        0.2,
		},
		{
			"id":          "ingredient-uuid-9",
			"photo":       "https://example.com/soy_sauce.jpg",
			"name":        "Soy Sauce",
			"measure":     "ml",
			"description": "Savory soy sauce.",
			"calories":    53.0,
			"proteins":    8.0,
			"water":       67.0,
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
