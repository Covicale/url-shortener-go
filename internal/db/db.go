package db

import (
	"database/sql"

	"github.com/covicale/url-shortener-go/internal/config"

	_ "github.com/lib/pq"
)

func Connect(databaseConfig config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseConfig.GetDSN())

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
