package utils

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func ParseBodyToJson(r *http.Request, payload interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Printf("Error: %v, Message: %s, StatusCode: %d\n", err, message, statusCode)

	errorMessage, _ := json.Marshal(map[string]string{"error": message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, writeErr := w.Write(errorMessage); writeErr != nil {
		log.Printf("Error writing the response: %v\n", writeErr)
	}
}

func GenerateRandomShortURL() string {
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shortUrl := make([]byte, 7)
	for i := range shortUrl {
		shortUrl[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortUrl)
}
