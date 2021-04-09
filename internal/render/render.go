package render

import (
	"bytes"
	"github.com/cknolla/go-webserver/internal/config"
	"github.com/cknolla/go-webserver/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(r)
	return templateData
}

func RenderTemplate(
	w http.ResponseWriter,
	r *http.Request,
	templateName string,
	data *models.TemplateData,
	) {
	var (
		templateCache map[string]*template.Template
		err error
	)

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = GetTemplateCache()
		if err != nil {
			log.Fatalln("Error getting template cache", err)
		}
	}

	tmpl, ok := templateCache[templateName]
	if !ok {
		log.Fatalln("Could not find template", templateName)
	}

	buffer := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	err = tmpl.Execute(buffer, data)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template buffer", err)
	}
}

// GetTemplateCache returns a template cache as a map
func GetTemplateCache() (map[string]*template.Template,error) {
	templateCache := map[string]*template.Template{}
	pagePaths, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, pagePath := range pagePaths {
		name := filepath.Base(pagePath)
		templateSet, err := template.
			New(name).
			Funcs(functions).
			ParseFiles(pagePath)
		if err != nil {
			return templateCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[name] = templateSet
	}

	return templateCache, nil
}
