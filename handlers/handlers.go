package handlers

import (
	"net/http"

	"github.com/orange432/mono-monero/models"
	"github.com/orange432/mono-monero/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html", &models.TemplateData{})
}
