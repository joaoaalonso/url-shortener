package repositories

import (
	"errors"

	"github.com/joaoaalonso/url-shortener/entities"
)

// URLRepository is a interface to data persistance
type URLRepository struct{}

var urls []entities.URL

// GetFromAlias find a url entity from alias
func (urlRepo *URLRepository) GetFromAlias(alias string) (entities.URL, error) {
	for _, url := range urls {
		if url.Alias == alias {
			return url, nil
		}
	}

	return entities.URL{}, errors.New("URL not found")
}

// Create new url
func (urlRepo *URLRepository) Create(url entities.URL) {
	urls = append(urls, url)
}
