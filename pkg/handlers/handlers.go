package handlers

import (
	"net/http"

	"github.com/thien0001/cmd/web/pkg/config"
	"github.com/thien0001/cmd/web/pkg/models"
	"github.com/thien0001/cmd/web/pkg/render"
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

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
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
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
