package routes

import (
	"github.com/gin-gonic/gin"
	"qr_url_service/internal/routes/v1"
	"qr_url_service/internal/handlers"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api")

	v1.RegisterRoutes(api.Group("/v1"))

	return router
}