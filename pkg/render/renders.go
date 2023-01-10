package render

import (
	"bytes"
	"github/dperkins/bookings/pkg/config"
	"github/dperkins/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// create sets the config for template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	//if use cache is true use that if not, create a new one each request
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from template app cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buffer, td)

	//this is to check if theres something wrong with the value in the map
	if err != nil {
		log.Println(err)
	}
	//render template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

const templatesPath string = "./templates/*.page.tmpl"
const layoutsPath string = "./templates/*.layout.tmpl"

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all files *.page.tmpl

	pages, err := filepath.Glob(templatesPath)

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//gives last piece of the path
		fileName := filepath.Base(page)
		//give the template a name
		ts, err := template.New(fileName).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		//filepath.glob returns slice
		layouts, err := filepath.Glob(layoutsPath)
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(layoutsPath)
		}

		myCache[fileName] = ts
	}
	return myCache, nil
}
