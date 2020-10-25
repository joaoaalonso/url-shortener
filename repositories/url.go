package repositories

import "github.com/joaoaalonso/url-shortener/entities"

// URLRepository interface
type URLRepository interface {
	GetFromAlias(alias string) (entities.URL, error)
	Create(url entities.URL) error
}
