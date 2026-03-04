package v1

import (
	controller "qr_url_service/internal/controllers/v1"
	"qr_url_service/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {

	urlService := services.NewURLService()

	urlController := controller.NewURLController(urlService)
	url := rg.Group("/url")
	{
		url.POST("/shorten", urlController.CreateShortURL)
	}
}
