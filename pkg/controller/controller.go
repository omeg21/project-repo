package controller

import (
	"net/http"

	"github.com/omeg21/project-repo/pkg/config"
	"github.com/omeg21/project-repo/pkg/models"
	"github.com/omeg21/project-repo/pkg/render"
)

// Repo the repository used by the controller
var Repo *Repository

// Repository is the repositry type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewController sets the repository for the handlers
func NewController(r *Repository) {
	Repo = r
}

// Home page controller
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteip)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there."

	remoteip := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteip

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// func New(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "new.page.html")
// }
