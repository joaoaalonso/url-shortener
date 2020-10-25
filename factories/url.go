package factories

import (
	"github.com/joaoaalonso/url-shortener/controllers"
	"github.com/joaoaalonso/url-shortener/repositories"
	"github.com/joaoaalonso/url-shortener/services"
)

// CreateURLController return a new instance of URLController
func CreateURLController() controllers.URLController {
	urlService := CreateURLService()
	return controllers.URLController{URLService: urlService}

}

// CreateURLService return a new instance of URLService
func CreateURLService() services.URLService {
	urlRepo := CreateURLRepository()
	return services.URLService{URLRepo: urlRepo}
}

// CreateURLRepository return a new instance of URLRepository
func CreateURLRepository() repositories.URLRepository {
	return repositories.URLRepository{}
}
