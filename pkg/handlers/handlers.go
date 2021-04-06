package handlers

import (
	"github.com/cknolla/go-webserver/pkg/config"
	"github.com/cknolla/go-webserver/pkg/models"
	"github.com/cknolla/go-webserver/pkg/render"
	"net/http"
)



// Repo is used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}


// NewRepo creates a new repository
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
