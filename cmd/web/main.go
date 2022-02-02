package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olumidesan/bookings/pkg/config"
	"github.com/olumidesan/bookings/pkg/handlers"
	"github.com/olumidesan/bookings/pkg/render"
)

const portNumber = ":3000"

// main is the entry to the weba app
func main() {
	// Create variable for app config
	var app config.AppConfig

	// Create the template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Config defaults
	app.TemplateCache = tc
	app.UseCache = true

	// Associate application context with route handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Associate application context with templates
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// Create a server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting HTTP Server...")
	err = srv.ListenAndServe()

	// Should be ideally unreachable
	log.Fatal("An error occured: ", err)
}
