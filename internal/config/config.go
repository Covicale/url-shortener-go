package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB               *DBConfig
	Port             string
	JWT_SECRET       string
	SHORTENER_DOMAIN string
}

type DBConfig struct {
	User     string
	Password string
	Name     string
	Address  string
	Port     string
}

var Env = initConfig()

func (db *DBConfig) GetDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		db.User, db.Password, db.Address, db.Port, db.Name)
}

func initConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file was found. Default variables will loaded")
	}

	dbConfig := DBConfig{
		User:     setConfigParam("DB_USER", "admin"),
		Password: setConfigParam("DB_PASSWORD", "admin"),
		Address:  setConfigParam("DB_ADDRESS", "localhost"),
		Name:     setConfigParam("DB_NAME", "postgres"),
		Port:     setConfigParam("DB_PORT", "5432"),
	}

	return &Config{
		DB:               &dbConfig,
		Port:             setConfigParam("PORT", "8000"),
		JWT_SECRET:       setConfigParam("JWT_SECRET", "secret"),
		SHORTENER_DOMAIN: setConfigParam("SHORTENER_DOMAIN", "http://localhost:8080"),
	}
}

func setConfigParam(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	return value
}
