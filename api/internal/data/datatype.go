package data

type URL struct {
	ShortUrl  string `json:"short_url"`
	ShortCode string `json:"short_code"`
}

type Envelope map[string]any
