package repositories

import (
	"database/sql"

	"github.com/covicale/url-shortener-go/internal/api/models"
)

type UserRepositoryInterface interface {
	CreateUser(models.User) error
	GetUserByEmail(string) (models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repository *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)"
	_, err := repository.db.Exec(query, user.Id, user.Username, user.Password, user.Email)
	return err
}

func (repository *UserRepository) GetUserByEmail(email string) (models.User, error) {
	query := "SELECT id, username, email, password, created_at FROM users WHERE email = $1"
	var user models.User
	if err := repository.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt); err != nil {
		return models.User{}, err
	}
	return user, nil
}
