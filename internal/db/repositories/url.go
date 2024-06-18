package repositories

import (
	"database/sql"
	"log"

	"github.com/covicale/url-shortener-go/internal/api/models"
)

type URLRepositoryInterface interface {
	CreateURL(*models.URL) error
	CreateRandomURL(*models.URL) error
	GetURLsByOwnerID(*models.URL) error
	GetLongURLByShortURL(string) (string, error)
	ExistsLongURLForUser(string, string) bool
	DeleteByShortURLAndOwnerID(string, string) (bool, error)
}

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (urlRepository *URLRepository) ExistsLongURLForUser(longURL string, userId string) bool {
	query := "SELECT long_url from urls WHERE long_url = $1 AND owner_id = $2"
	err := urlRepository.db.QueryRow(query, longURL, userId).Scan(&longURL)
	return err == sql.ErrNoRows
}

func (urlRepository *URLRepository) GetURLsByOwnerID(userId string) ([]models.URL, error) {
	query := "SELECT id, short_url, long_url, owner_id from urls WHERE owner_id = $1"
	var urls []models.URL
	rows, err := urlRepository.db.Query(query, userId)
	if err != nil {
		return urls, err
	}
	for rows.Next() {
		var url models.URL
		if err := rows.Scan(
			&url.Id,
			&url.ShortURL,
			&url.LongURL,
			&url.OwnerId,
		); err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (urlRepository *URLRepository) CreateURL(url *models.URL) error {
	query := "INSERT INTO urls (owner_id, short_url, long_url) VALUES($1, $2, $3)"
	_, err := urlRepository.db.Exec(query, url.OwnerId, url.ShortURL, url.LongURL)
	if err != nil {
		log.Println("An error ocurred creating the URL", err)
	}
	return nil
}

func (urlRepository *URLRepository) GetLongURLByShortURL(shortURL string) (string, error) {
	query := "SELECT long_url FROM urls WHERE short_url = $1"
	var longURL string
	if err := urlRepository.db.QueryRow(query, shortURL).Scan(&longURL); err != nil {
		log.Println("An error ocurred getting the long URL", err)
		return "", err
	}
	return longURL, nil
}

func (urlRepository *URLRepository) DeleteByShortURLAndOwnerID(shortUrl string, ownerId string) (bool, error) {
	query := "DELETE FROM urls WHERE short_url = $1 AND owner_id = $2"
	result, err := urlRepository.db.Exec(query, shortUrl, ownerId)
	if err != nil {
		return false, err
	}
	rowsChanged, _ := result.RowsAffected()
	return rowsChanged == 1, nil
}
