package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/covicale/url-shortener-go/internal/api/models"
	"github.com/covicale/url-shortener-go/internal/api/types"
	"github.com/covicale/url-shortener-go/internal/api/utils"
	"github.com/covicale/url-shortener-go/internal/db/repositories"
)

type AuthHandler struct {
	userRepository *repositories.UserRepository
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (h *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	jwtDeleted := utils.DeleteJWTCookie()
	http.SetCookie(w, jwtDeleted)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var rBody types.RegisterUserRequest
	if err := utils.ParseBodyToJson(r, &rBody); err != nil {
		utils.WriteError(w, err, "Body not valid", http.StatusBadRequest)
		return
	}

	if err := utils.CheckCredentialsOnRegister(&rBody); err != nil {
		utils.WriteError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rBody.Password), 4)
	if err != nil {
		utils.WriteError(w, err, "Internal Server Error", http.StatusInternalServerError)
	}

	user := models.NewUser(rBody.Username, string(hashedPassword), rBody.Email)

	userClaims := utils.NewUserJWTClaims(user)
	jwtCookie, err := utils.CreateJWTCookie(userClaims)
	if err != nil {
		utils.WriteError(w, err, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := h.userRepository.CreateUser(user); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			utils.WriteError(w, err, "This username/email already exists.", http.StatusBadRequest)
		} else {
			utils.WriteError(w, err, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, jwtCookie)
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var rBody types.LoginUserRequest
	if err := utils.ParseBodyToJson(r, &rBody); err != nil {
		utils.WriteError(w, err, "Body not valid", http.StatusBadRequest)
		return
	}

	user, err := h.userRepository.GetUserByEmail(rBody.Email)
	if err != nil {
		utils.WriteError(w, err, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rBody.Password))
	if err != nil {
		utils.WriteError(w, err, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	// Successfully logged in
	userClaims := utils.NewUserJWTClaims(&user)
	jwtCookie, err := utils.CreateJWTCookie(userClaims)
	if err != nil {
		utils.WriteError(w, err, "Internal server error", http.StatusInternalServerError)
	}

	http.SetCookie(w, jwtCookie)
	w.Write([]byte("User successfully logged in."))
}
