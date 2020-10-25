package repositories

import (
	"errors"

	"github.com/joaoaalonso/url-shortener/entities"
)

// URLMemoryRepository is a interface to data persistance
type URLMemoryRepository struct{}

var urls []entities.URL

// GetFromAlias find a url entity from alias
func (urlRepo *URLMemoryRepository) GetFromAlias(alias string) (entities.URL, error) {
	for _, url := range urls {
		if url.Alias == alias {
			return url, nil
		}
	}

	return entities.URL{}, errors.New("URL not found")
}

// Create new url in memory
func (urlRepo *URLMemoryRepository) Create(url entities.URL) error {
	urls = append(urls, url)

	return nil
}
