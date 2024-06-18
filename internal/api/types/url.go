package types

import "github.com/google/uuid"

type CreateURLRequest struct {
	URL    string    `json:"url"`
	UserId uuid.UUID `json:"userId"`
}

type DeleteURLRequest struct {
	ShortURL string `json:"shortUrl"`
}
