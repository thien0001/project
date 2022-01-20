package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"

	"github.com/thien0001/cmd/web/internal/config"
	"github.com/thien0001/cmd/web/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaulData adds data for all temlaptes
func AddDefaulData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

//Render Template renders templates using html/templates
// func RenderTemplate(w http.ResponseWriter, html string) {
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {

		// get the template cache from the app config

		tc = app.TemplateCache
	} else {

		tc, _ = CreateTemplateCache()
	}

	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	t, ok := tc[html]

	if !ok {
		// log.Fatal(err)
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaulData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		if err != nil {
			fmt.Println("Error writting template to browser", err)
		}
	}

}

// CreatetemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {

		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts

	}
	return myCache, nil

}
