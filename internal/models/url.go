package models

type Url struct {
	ID        int    `json:"id"`
	ShortUrl  string `json:"short_url"`
	LongUrl   string `json:"long_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}