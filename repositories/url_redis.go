package repositories

import (
	"context"
	"errors"

	"github.com/joaoaalonso/url-shortener/configs"

	"github.com/joaoaalonso/url-shortener/entities"
)

// URLRedisRepository is a interface to data persistance
type URLRedisRepository struct{}

var ctx = context.Background()

// GetFromAlias find a url entity from alias
func (urlRepo *URLRedisRepository) GetFromAlias(alias string) (entities.URL, error) {
	client := configs.GetRedisConnection()
	longURL, err := client.Get(ctx, alias).Result()

	if err != nil {
		return entities.URL{}, errors.New("URL not found")
	}

	return entities.URL{LongURL: longURL, Alias: alias}, nil
}

// Create new url in redis
func (urlRepo *URLRedisRepository) Create(url entities.URL) error {
	client := configs.GetRedisConnection()
	status := client.Set(ctx, url.Alias, url.LongURL, 0)

	return status.Err()
}
