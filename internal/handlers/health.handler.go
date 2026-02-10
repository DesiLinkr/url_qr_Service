package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func HealthCheck(context *gin.Context) {

	context.JSON(200, gin.H{
		"status":    "ok",
		"uptime":    time.Since(startTime).Seconds(),
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
