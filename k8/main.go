package main

import (
	"1/controllers"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	c, _ := elasticsearch.NewDefaultClient()

	controller := controllers.NewApi(c)

	// Routes
	e.GET("/", controller.Hello)

	e.GET("/healthCheck", controller.Health)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
