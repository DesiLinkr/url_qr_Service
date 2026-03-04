package v1

type ShortURLRequest struct {
	URL string `json:"url" binding:"required"`
}
