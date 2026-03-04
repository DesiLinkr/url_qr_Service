package services

type URLService struct {
}

func NewURLService() *URLService {
	return &URLService{}
}

func (s *URLService) GenerateShortURL(url string) string {

	return "abc123"
}
