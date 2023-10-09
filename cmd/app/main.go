package main

import (
	"github.com/ensamblaTec/learning/week2/problema4/database"
	"github.com/ensamblaTec/learning/week2/problema4/pkg/handler"
	"github.com/ensamblaTec/learning/week2/problema4/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init database
	database.Initialize() // if exists
	handler.InitializeProducts()
	// end database
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	e.Static("/", "web/static")

	tmpl := models.Init()

	e.Renderer = tmpl

	e.GET("/", handler.RunApp)

	e.POST("/products", handler.RegisterProduct)

	e.DELETE("/products", handler.DeleteProduct)

	e.Logger.Fatal(e.Start(":80"))
}
