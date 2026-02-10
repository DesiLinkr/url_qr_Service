package routes

import (
	"qr_url_service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRegister(healthRouter *gin.RouterGroup) {

	healthRouter.GET("/", handlers.HealthCheck)

}
