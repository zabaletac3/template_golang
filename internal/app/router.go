package app

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eren_dev/go_server/internal/app/middleware"
)

func registerRoutes(router *gin.Engine) {
	// Power off internal Gin logs
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	// Custom middlewares
	router.Use(middleware.SlogLogger())
	router.Use(middleware.SlogRecovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API is running"})
	})

	// api := router.Group("/api")

	// users.Register(api)
}
