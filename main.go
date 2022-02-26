package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/orange432/mono-monero/cache"
	"github.com/orange432/mono-monero/captcha"
	"github.com/orange432/mono-monero/config"
	"github.com/orange432/mono-monero/handlers"
	"github.com/orange432/mono-monero/render"
)

var PORT = ":3000"

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomePage).Methods("GET")
	r.HandleFunc("/login", handlers.LoginPage).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterPage).Methods("GET")
	fileServer := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
	return r
}

func Init() error {
	// Load up the templates
	var appCache cache.AppCache
	config.InitConfig("./config.json")
	tCache, err := render.CreateTemplateCache()
	appCache.LastGotTemplates = time.Now().Unix()
	if err != nil {
		log.Fatal("Couldn't load template cache")
		return err
	}
	appCache.Templates = tCache

	render.NewTemplates(&appCache)
	captcha.LoadCache(&appCache)
	return nil
}

func main() {
	// Initializ
	err := Init()
	if err != nil {
		log.Fatal(err)
	}

	// Get routes
	r := Routes()
	fmt.Println(fmt.Sprintf("🚀 Running at http://localhost%s", PORT))
	log.Fatal(http.ListenAndServe(PORT, r))
}
