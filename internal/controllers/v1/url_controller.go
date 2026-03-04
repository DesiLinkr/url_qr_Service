package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	schema_v1 "qr_url_service/internal/schema/v1"
	"qr_url_service/internal/services"
)

type URLController struct {
	urlService *services.URLService
}

func NewURLController(urlService *services.URLService) *URLController {
	return &URLController{
		urlService: urlService,
	}
}

func (u *URLController) CreateShortURL(c *gin.Context) {

	var req schema_v1.ShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortCode := u.urlService.GenerateShortURL(req.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_code": shortCode,
	})

}
