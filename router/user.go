package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (r *Router) LoginHandler(ctx echo.Context) (err error) {

	// Simulate successful login (replace with actual authentication logic)
	userMap := map[string]interface{}{
		"id":        "your-generated-uuid",
		"firstname": "daniel",                                                                                                                           // Assuming name is the user's first name
		"lastname":  "Doe",                                                                                                                              // Placeholder, replace with actual last name
		"photo":     "https://media.istockphoto.com/id/639805094/photo/happy-man.jpg?s=612x612&w=0&k=20&c=REx0Usczge4a0soQvp7fQgGCcFMHeORGUTpOIPW-IYA=", // Placeholder, replace with actual photo URL (optional)
		"email":     "danielvieira2828@gmail.com",
		"token":     "your-generated-token", // Replace with actual token
		"height":    175.0,                  // Placeholder, replace with actual height (optional)
		"weight":    70.0,                   // Placeholder, replace with actual weight (optional)
		"imc":       10,                     // Replace with actual IMC calculation
		"sex":       "male",                 // Placeholder, replace with actual sex (optional)
		"goals": map[string]bool{
			"gainMuscularMass": false, // Placeholder, adjust based on user's goals (optional)
			"maintainHealth":   true,  // Placeholder, adjust based on user's goals (optional)
			"loseWeight":       false, // Placeholder, adjust based on user's goals (optional)
		},
		"biotype":   "endomorph", // Placeholder, replace with actual biotype (optional)
		"createdAt": time.Now().Format(time.RFC3339Nano),
		"updatedAt": time.Now().Format(time.RFC3339Nano),
	}

	// Encode user model to JSON
	userJSON, err := json.Marshal(userMap)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error encoding user to JSON: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, userJSON)
}

func (r *Router) RegisterHandler(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "Você deve retornar o UUID gerado pelo consumer")
}
