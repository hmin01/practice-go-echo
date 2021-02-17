package routes

import (
	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// Handler
	handler "practice-go-echo/handlers"
)

func Router() *echo.Echo {
	// Create echo
	app := echo.New()
	// Set middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// Set static
	app.Static("/assets", "static")
	// Set Health check api
	app.GET("/health", handler.HealthCheck)

	return app
}