package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func init() {

	// Init logging
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// Check if required env vars are present
	requiredEnvVars := []string{"DB_USER", "DB_PASS",
		"DB_HOST", "DB_PORT", "DB_NAME"}
	for _, envar := range requiredEnvVars {
		if val, ok := os.LookupEnv(envar); !ok || val == "" {
			zap.L().Fatal(fmt.Sprintf("%s is missing", envar))
		}
	}

}

func Run() {

	// Create a new instance of Echo framework
	e := echo.New()

	// Init in-memory datastore
	initDatastore()

	// Init bookstore API routes
	initRoutes(e)

	// Health check
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	// Version info
	e.GET("/version", func(ctx echo.Context) error {
		resp := make(map[string]interface{})
		resp["version"] = os.Getenv("LATEST_GIT_TAG")
		return ctx.JSON(http.StatusOK, resp)
	})

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Run HTTP server
	if err := e.Start(":8080"); err != nil {
		zap.L().Fatal(err.Error())
	}
}
