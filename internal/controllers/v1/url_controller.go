package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

	// read request
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	shortURL := "abc123"

	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})

}
