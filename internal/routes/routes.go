package routes

import (
	"vulnlabz/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	rootHandler := handlers.NewRootHandler()
	healthHandler := handlers.NewHealthHandler()

	api := router.Group("/api/v1")
	{
		api.GET("/", rootHandler.Root)
		api.GET("/health", healthHandler.HealthCheck)
	}

	router.GET("/", rootHandler.Root)
	router.GET("/health", healthHandler.HealthCheck)
}
