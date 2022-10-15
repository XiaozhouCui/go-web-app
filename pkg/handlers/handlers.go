package handlers

import (
	"gowebapp/pkg/config"
	"gowebapp/pkg/render"
	"net/http"
)

// repository pattern is used for handlers

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers (from app config)
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// m is the receiver, giving Home access to everything in the repository
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// to refer to this func, use "handlers.Repo.About"
	render.RenderTemplate(w, "about.page.tmpl")
}
