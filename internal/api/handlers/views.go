package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path"

	"github.com/covicale/url-shortener-go/internal/api/models"
	"github.com/covicale/url-shortener-go/internal/api/utils"
	"github.com/covicale/url-shortener-go/internal/db/repositories"
)

type ViewHandler struct {
	userRepository *repositories.UserRepository
	urlRepository  *repositories.URLRepository
}

func NewViewHandler(db *sql.DB) *ViewHandler {
	return &ViewHandler{
		userRepository: repositories.NewUserRepository(db),
		urlRepository:  repositories.NewURLRepository(db),
	}
}

func (handler *ViewHandler) ServeFavicon(w http.ResponseWriter, r *http.Request) {
	faviconPath := path.Join("views", "favicon.svg")
	http.ServeFile(w, r, faviconPath)
}

func (handler *ViewHandler) RedirectToRealURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("url")
	if shortURL == "" {
		return
	}
	realURL, err := handler.urlRepository.GetLongURLByShortURL(shortURL)
	if err != nil {
		utils.WriteError(w, err, "An error ocurred trying to retrieve the real url.", http.StatusBadGateway)
		return
	}
	http.Redirect(w, r, realURL, http.StatusFound)
}

func (handler *ViewHandler) RenderHome(w http.ResponseWriter, r *http.Request) {
	userInfo, _ := utils.GetUserInfoFromRequest(r)

	urls, _ := handler.urlRepository.GetURLsByOwnerID(userInfo.UserID.String())
	pageData := models.NewPageDataHome(&userInfo, &urls)

	filePath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		utils.WriteError(w, err, "Error parsing the index.html file", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		utils.WriteError(w, err, "Error rendering the page", http.StatusInternalServerError)
	}
}

func (handler *ViewHandler) RenderLogin(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("views", "auth", "login.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		utils.WriteError(w, err, "Error parsing the index.html file", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		utils.WriteError(w, err, "Error rendering the page", http.StatusInternalServerError)
	}
}

func (handler *ViewHandler) RenderRegister(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("views", "auth", "register.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		utils.WriteError(w, err, "Error parsing the register.html file", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		utils.WriteError(w, err, "Error rendering the page", http.StatusInternalServerError)
	}
}
