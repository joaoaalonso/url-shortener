package services

import (
	"github.com/joaoaalonso/url-shortener/entities"
	"github.com/joaoaalonso/url-shortener/repositories"
	uuid "github.com/satori/go.uuid"
)

// URLService contains all url use cases
type URLService struct {
	URLRepo repositories.URLRepository
}

// GetURLFromAlias returns a redirect url from alias
func (urlService *URLService) GetURLFromAlias(alias string) string {
	url, err := urlService.URLRepo.GetFromAlias(alias)
	if err != nil {
		return "/"
	}

	return url.LongURL
}

// GenerateAlias returns a random unique alias
func (urlService *URLService) GenerateAlias() string {
	alias := uuid.NewV4()
	return alias.String()[0:7]
}

// Create new url
func (urlService *URLService) Create(longURL string, alias string) (entities.URL, error) {
	if alias == "" {
		alias = urlService.GenerateAlias()
	}

	url := entities.URL{
		LongURL: longURL,
		Alias:   alias,
	}

	urlService.URLRepo.Create(url)

	return url, nil
}
