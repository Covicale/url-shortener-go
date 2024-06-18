package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/types"
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
	// Log the original error with more context
	log.Printf("Error: %v, Message: %s, StatusCode: %d\n", err, message, statusCode)

	errorMessage, _ := json.Marshal(map[string]string{"error": message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, writeErr := w.Write(errorMessage); writeErr != nil {
		log.Printf("Error writing the response: %v\n", writeErr)
	}
}

func GetUserInfoFromRequest(r *http.Request) (types.UserJWT, error) {
	var userInfo types.UserJWT
	userInfoJson := r.Header.Get("UserInfo")
	if userInfoJson != "" {
		json.Unmarshal([]byte(userInfoJson), &userInfo)
		return userInfo, nil
	}
	return types.UserJWT{}, errors.New("UserInfo doesn't exist in header")
}

func GenerateRandomShortURL() string {
	randNumber := rand.Intn(int(math.Pow(62, 7)))
	return EncodeToBase62(randNumber)
}
