package main

import (
	"github.com/ensamblaTec/learning/week2/problema4/pkg/handler"
	"github.com/ensamblaTec/learning/week2/problema4/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	c := echo.New()

	c.Use(middleware.Logger())

	tmpl := models.Init()

	c.Renderer = tmpl

	c.GET("/", handler.RunLogin)

	c.POST("/add-product", handler.RegisterProduct)

	c.Logger.Fatal(c.Start(":80"))
}
