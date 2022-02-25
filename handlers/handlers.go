package handlers

import (
	"net/http"

	"github.com/orange432/mono-monero/models"
	"github.com/orange432/mono-monero/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html", &models.TemplateData{})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.html", &models.TemplateData{})
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.html", &models.TemplateData{})
}
