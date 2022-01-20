package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/thien0001/cmd/web/internal/config"
	"github.com/thien0001/cmd/web/internal/models"
	"github.com/thien0001/cmd/web/internal/render"
)

// // TemplateData holds data sent from handlers package
// type TemplateData struct {
// 	StringMap map[string]string
// 	InMap     map[string]int
// 	FloatMap  map[string]float32
// 	Data      map[string]interface{}
// 	CSRFToken string
// 	Flash     string
// 	Warning   string
// 	Error     string
// }

// Repo the repository used by the handlers-
var Repo *Respository

// Respository is the repository type

type Respository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandlers(r *Respository) {
	Repo = r
}

// Home is the home page handler

// func Home(w http.ResponseWriter, r *http.Request) {
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

//Sprintf in ra %d
// func About(w http.ResponseWriter, r *http.Request) {

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	// stringMap["remote_ip"] = remoteIP

	//perform some logic
	// send the data the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation renders the make a verservation page and displays form
func (m *Respository) Rerservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

// Generals render the room
func (m *Respository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors render the room
func (m *Respository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

//Availability
func (m *Respository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availabilly.page.html", &models.TemplateData{})
}

// PostAvailability
func (m *Respository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON
func (m *Respository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Contet-Type", "application/json")
	w.Write(out)
}

// Contact
func (m *Respository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// Makereservation
func (m *Respository) Makereservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}
