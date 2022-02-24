package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/orange432/mono-monero/cache"
	"github.com/orange432/mono-monero/models"
)

var appCache *cache.AppCache

var PATH_TO_PAGES = "./templates/pages"
var PATH_TO_LAYOUTS = "./templates/layouts"

// NewTemplates loads in the cache variable
func NewTemplates(c *cache.AppCache) {
	appCache = c
}

// RenderTemplate renders a specific template from the template cache.
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var templates map[string]*template.Template

	templates = appCache.Templates

	t, ok := templates[tmpl]
	if !ok {
		log.Fatal("Couldn't load templates.")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", w)
	}
}

// Loads up all the templates into memory
func CreateTemplateCache() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", PATH_TO_PAGES))
	if err != nil {
		return tCache, err
	}

	for _, page := range pages {
		// Load normal page content
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tCache, err
		}

		// Load layouts
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.html", PATH_TO_LAYOUTS))
		if err != nil {
			return tCache, err
		}

		// Layouts found, add to template string
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.html", PATH_TO_LAYOUTS))
			if err != nil {
				return tCache, err
			}
		}
		tCache[name] = ts
	}
	// Gone through each page and template, return success
	return tCache, nil
}
