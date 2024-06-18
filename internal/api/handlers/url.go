package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/models"
	"github.com/covicale/url-shortener-go/internal/api/types"
	"github.com/covicale/url-shortener-go/internal/api/utils"
	"github.com/covicale/url-shortener-go/internal/db/repositories"
)

type URLHandler struct {
	urlRepository *repositories.URLRepository
}

func NewURLHandler(db *sql.DB) *URLHandler {
	return &URLHandler{
		urlRepository: repositories.NewURLRepository(db),
	}
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	userInfo, err := utils.GetUserInfoFromRequest(r)
	if err != nil {
		utils.WriteError(w, nil, "Need to be logged in", http.StatusUnauthorized)
		return
	}

	var rBody types.CreateURLRequest
	if err := utils.ParseBodyToJson(r, &rBody); err != nil {
		utils.WriteError(w, err, "Body not valid", http.StatusBadRequest)
		return
	}

	// Check if the url is valid
	if !utils.IsURLValid(rBody.URL) {
		utils.WriteError(w, errors.New("URL not valid"), "URL not valid", http.StatusBadRequest)
		return
	}

	// Retrieve the user info with the cookies and jwt
	urlShorten := models.NewURL(utils.GenerateRandomShortURL(), rBody.URL, userInfo.UserID)

	// Check if already exists, and if exists, return the short url
	if exists := h.urlRepository.ExistsLongURLForUser(urlShorten.LongURL, userInfo.UserID.String()); !exists {
		utils.WriteError(w, nil, "You already have reduce this url", http.StatusConflict)
		return
	}

	if err := h.urlRepository.CreateURL(urlShorten); err != nil {
		utils.WriteError(w, err, "Error writing in the database", http.StatusInternalServerError)
		return
	}

	urlMap := map[string]string{
		"shortURL": urlShorten.ShortURL,
	}
	response, err := json.Marshal(urlMap)
	if err != nil {
		utils.WriteError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *URLHandler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	userInfo, err := utils.GetUserInfoFromRequest(r)
	if err != nil {
		utils.WriteError(w, nil, "Need to be logged in", http.StatusUnauthorized)
		return
	}

	shortURL := r.PathValue("shortURL")

	if deleted, err := h.urlRepository.DeleteByShortURLAndOwnerID(shortURL, userInfo.UserID.String()); err != nil || !deleted {
		utils.WriteError(w, nil, "Error removing the url from the database", http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
