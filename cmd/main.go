package main

import (
	"fmt"
	"qr_url_service/internal/config"
	"qr_url_service/internal/routes"
)

func main() {
	port := config.Port
	router := routes.SetupRouter()
	fmt.Println("Server running on :", port)
	router.Run(":" + port)
}
