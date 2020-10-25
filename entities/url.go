package entities

// URL contains shorten url informations
type URL struct {
	Alias   string `json:"alias"`
	LongURL string `json:"long_url"`
}