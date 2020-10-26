package services

import (
	"errors"
	"regexp"

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
func (urlService *URLService) GenerateAlias(countOptional ...int) (string, error) {
	count := 0
	if len(countOptional) > 0 {
		count = countOptional[0]
	}

	if count >= 5 {
		return "", errors.New("Can't generate unique alias")
	}

	fullAlias := uuid.NewV4()
	alias := fullAlias.String()[0:7]
	var err error

	if urlService.aliasExists(alias) {
		alias, err = urlService.GenerateAlias(count + 1)
	}

	return alias, err
}

// Create new url
func (urlService *URLService) Create(longURL string, alias string) (entities.URL, error) {
	if alias == "" {
		newAlias, err := urlService.GenerateAlias()
		if err != nil {
			return entities.URL{}, err
		}
		alias = newAlias
	} else if urlService.aliasExists(alias) {
		return entities.URL{}, errors.New("Alias alredy in use")
	} else if !urlService.aliasIsValid(alias) {
		return entities.URL{}, errors.New("Not allowed alias")
	}

	url := entities.URL{
		LongURL: longURL,
		Alias:   alias,
	}

	err := urlService.URLRepo.Create(url)

	return url, err
}

func (urlService *URLService) aliasExists(alias string) bool {
	_, err := urlService.URLRepo.GetFromAlias(alias)
	return err == nil
}

func (urlService *URLService) aliasIsValid(alias string) bool {
	regex, _ := regexp.Compile("^[A-Za-z0-9_-]+$")

	return regex.Match([]byte(alias))
}
