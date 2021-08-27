package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Add your handlers here

// exampleHandler returns 200 and plain text
func exampleHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello from your service")
}