package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/omeg21/project-repo/internal/config"
	"github.com/omeg21/project-repo/internal/models"
	"github.com/omeg21/project-repo/internal/render"
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
	render.RenderTemplate(w, r,  "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there."

	remoteip := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteip

	//send the data to the template
	render.RenderTemplate(w,r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// func New(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "new.page.html")
// }

// Renders the Primo roompage
func (m *Repository) Primo(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r,"generals.page.html",&models.TemplateData{})
	
}
// Renders the Jojo roompage
func (m *Repository) Jojo(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r,"majors.page.html",&models.TemplateData{})
	
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r,"searchAvailability.page.html",&models.TemplateData{})
	
}
// Post Availability form
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end dae is %s", start, end)))
	
}
type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}
// Post JSON response form
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK :false ,
		Message: "Available!",
	}

	out,err := json.MarshalIndent(resp,"","     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-type","appllication/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r,"contact.page.html",&models.TemplateData{})
	
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r,"makeReservation.page.html",&models.TemplateData{})
	
}