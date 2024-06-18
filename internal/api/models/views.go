package models

import (
	"github.com/covicale/url-shortener-go/internal/api/types"
	"github.com/covicale/url-shortener-go/internal/config"
)

type PageDataHome struct {
	User   *types.UserJWT
	URLS   *[]URL
	Domain string
}

func NewPageDataHome(user *types.UserJWT, urls *[]URL) *PageDataHome {
	return &PageDataHome{
		User:   user,
		URLS:   urls,
		Domain: config.Env.SHORTENER_DOMAIN,
	}
}
