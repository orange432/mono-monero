package handlers

import (
	"net/http"

	"github.com/orange432/mono-monero/captcha"
	"github.com/orange432/mono-monero/models"
	"github.com/orange432/mono-monero/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html", &models.TemplateData{})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	img, hash := captcha.Generate()
	render.RenderTemplate(w, "login.html", &models.TemplateData{
		Captcha:     img,
		CaptchaHash: hash,
	})
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	img, hash := captcha.Generate()
	render.RenderTemplate(w, "register.html", &models.TemplateData{
		Captcha:     img,
		CaptchaHash: hash,
	})
}
