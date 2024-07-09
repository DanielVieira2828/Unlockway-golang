package server

import (
	"log"

	"github.com/DanielVieirass/Unlockway_golang/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() error {
	server := echo.New()
	server.Use(middleware.Recover())

	router := &router.Router{
		Server: server,
	}

	var user = server.Group("/user")
	var dishes = server.Group("/dishes")
	var routines = server.Group("/routines")
	var history = server.Group("/history")
	var notify = server.Group("/notify")

	user.POST("/login", router.LoginHandler)
	user.GET("/register", router.RegisterHandler)

	dishes.GET("/get", router.GetDishes)
	dishes.GET("/create", router.CreateDish)
	dishes.GET("/update", router.UpdateDish)
	dishes.GET("/delete", router.DeleteDish)
	dishes.GET("/get/ingredients", router.GetIngredients)

	routines.GET("/get", router.GetRoutines)
	routines.GET("/update", router.UpdateRoutine)
	routines.GET("/create", router.CreateRoutine)
	routines.GET("/delete", router.DeleteRoutine)

	history.GET("/get", router.GetHistory)
	history.GET("/update", router.UpdateHistory)
	history.GET("/delete", router.DeleteHistory)
	history.GET("/create", router.CreateHistory)

	notify.GET("/get", router.GetNotify)

	server.HideBanner = true
	server.HidePort = true

	log.Println("starting server...")

	return server.Start(":8080")
}
