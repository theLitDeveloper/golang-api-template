package service

import "github.com/labstack/echo/v4"

// initRoutes sets up all API routes
func initRoutes(e *echo.Echo) {
	e.GET("/", exampleHandler)
}
