package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRegister(apiRouter *gin.RouterGroup) {

	healthRouter := apiRouter.Group("/health")

	HealthRegister(healthRouter)

}
