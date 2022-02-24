package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orange432/mono-monero/cache"
	"github.com/orange432/mono-monero/handlers"
	"github.com/orange432/mono-monero/render"
)

var PORT = ":3000"

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomePage).Methods("GET")

	return r
}

func Init() error {
	// Load up the templates
	var appCache cache.AppCache

	tCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Couldn't load template cache")
		return err
	}
	appCache.Templates = tCache

	render.NewTemplates(&appCache)
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
