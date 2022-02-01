package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/olumidesan/bookings/pkg/config"
	"github.com/olumidesan/bookings/pkg/handlers"
	"github.com/olumidesan/bookings/pkg/render"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting HTTP Server...")
	_ = http.ListenAndServe(":3000", nil)
}
