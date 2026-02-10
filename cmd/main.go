package main

import (
	"fmt"
	"qr_url_service/internal/config"
	"qr_url_service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	port := config.Port
	apiRouter := app.Group("/api")
	routes.ApiRegister(apiRouter)
	fmt.Println("Server running on :", port)
	app.Run(":" + port)
}
