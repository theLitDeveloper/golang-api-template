package service

import "github.com/labstack/echo/v4"

// initRoutes sets up all routes for this API
func initRoutes(e *echo.Echo) {
	e.GET("/books", booksHandler)
	e.GET("/books/:id", getBookByIDHandler)
	e.POST("/books", addBookHandler)
	e.PATCH("/books/:id", updateBookHandler)
	e.DELETE("/books/:id", deleteBookHandler)
}
