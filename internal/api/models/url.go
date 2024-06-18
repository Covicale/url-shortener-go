package models

import (
	"github.com/google/uuid"
)

type URL struct {
	Id       int       `json:"id"`
	ShortURL string    `json:"short_url"`
	LongURL  string    `json:"long_url"`
	OwnerId  uuid.UUID `json:"owner_id"`
}

func NewURL(shortUrl, longURL string, ownerID uuid.UUID) *URL {
	return &URL{
		ShortURL: shortUrl,
		LongURL:  longURL,
		OwnerId:  ownerID,
	}
}
