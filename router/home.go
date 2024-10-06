package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Router) GetHomeData(ctx echo.Context) error {
	responseData := map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"meals":         2,
			"routines":      2,
			"notifications": 8,
			"weekCalories":  []interface{}{2805.05, 1032.05, 2100.05, 165.05, 4000.00, nil, nil},
		},
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func (r *Router) GetNutriHomeData(ctx echo.Context) error {
	responseData := map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"meals":         2,
			"routines":      2,
			"notifications": 8,
			"weekCalories":  []interface{}{2805.05, 1032.05, 2100.05, 165.05, 4000.00, nil, nil},
			"userList": []map[string]interface{}{
				{"id": "your-generated-uuid",
					"firstname": "Daniel",                                                                                                                           // Assuming name is the user's first name
					"lastname":  "Doe",                                                                                                                              // Placeholder, replace with actual last name
					"photo":     "https://media.istockphoto.com/id/639805094/photo/happy-man.jpg?s=612x612&w=0&k=20&c=REx0Usczge4a0soQvp7fQgGCcFMHeORGUTpOIPW-IYA=", // Placeholder, replace with actual photo URL (optional)
					"email":     "danielvieira2828@gmail.com",
					"token":     "your-generated-token", // Replace with actual token
					"height":    175.05,                 // Placeholder, replace with actual height (optional)
					"weight":    70.05,                  // Placeholder, replace with actual weight (optional)
					"imc":       10.05,                  // Replace with actual IMC calculation
					"sex":       "MALE",                 // Placeholder, replace with actual sex (optional)
					"goals": map[string]bool{
						"gainMuscularMass": false, // Placeholder, adjust based on user's goals (optional)
						"maintainHealth":   true,  // Placeholder, adjust based on user's goals (optional)
						"loseWeight":       false, // Placeholder, adjust based on user's goals (optional)
					},
					"biotype":   "ENDOMORPH", // Placeholder, replace with actual biotype (optional)
					"createdAt": time.Now().Format(time.RFC3339Nano),
					"updatedAt": time.Now().Format(time.RFC3339Nano),
				},
				{"id": "your-generated-uuid",
					"firstname": "Daniel",                                                                                                                           // Assuming name is the user's first name
					"lastname":  "Doe",                                                                                                                              // Placeholder, replace with actual last name
					"photo":     "https://media.istockphoto.com/id/639805094/photo/happy-man.jpg?s=612x612&w=0&k=20&c=REx0Usczge4a0soQvp7fQgGCcFMHeORGUTpOIPW-IYA=", // Placeholder, replace with actual photo URL (optional)
					"email":     "danielvieira2828@gmail.com",
					"token":     "your-generated-token", // Replace with actual token
					"height":    175.05,                 // Placeholder, replace with actual height (optional)
					"weight":    70.05,                  // Placeholder, replace with actual weight (optional)
					"imc":       10.05,                  // Replace with actual IMC calculation
					"sex":       "MALE",                 // Placeholder, replace with actual sex (optional)
					"goals": map[string]bool{
						"gainMuscularMass": false, // Placeholder, adjust based on user's goals (optional)
						"maintainHealth":   true,  // Placeholder, adjust based on user's goals (optional)
						"loseWeight":       false, // Placeholder, adjust based on user's goals (optional)
					},
					"biotype":   "ENDOMORPH", // Placeholder, replace with actual biotype (optional)
					"createdAt": time.Now().Format(time.RFC3339Nano),
					"updatedAt": time.Now().Format(time.RFC3339Nano),
				},
				{"id": "your-generated-uuid",
					"firstname": "Daniel",                                                                                                                           // Assuming name is the user's first name
					"lastname":  "Doe",                                                                                                                              // Placeholder, replace with actual last name
					"photo":     "https://media.istockphoto.com/id/639805094/photo/happy-man.jpg?s=612x612&w=0&k=20&c=REx0Usczge4a0soQvp7fQgGCcFMHeORGUTpOIPW-IYA=", // Placeholder, replace with actual photo URL (optional)
					"email":     "danielvieira2828@gmail.com",
					"token":     "your-generated-token", // Replace with actual token
					"height":    175.05,                 // Placeholder, replace with actual height (optional)
					"weight":    70.05,                  // Placeholder, replace with actual weight (optional)
					"imc":       10.05,                  // Replace with actual IMC calculation
					"sex":       "MALE",                 // Placeholder, replace with actual sex (optional)
					"goals": map[string]bool{
						"gainMuscularMass": false, // Placeholder, adjust based on user's goals (optional)
						"maintainHealth":   true,  // Placeholder, adjust based on user's goals (optional)
						"loseWeight":       false, // Placeholder, adjust based on user's goals (optional)
					},
					"biotype":   "ENDOMORPH", // Placeholder, replace with actual biotype (optional)
					"createdAt": time.Now().Format(time.RFC3339Nano),
					"updatedAt": time.Now().Format(time.RFC3339Nano),
				},
			},
		},
	}

	print(responseData)

	return ctx.JSON(http.StatusOK, responseData)
}
