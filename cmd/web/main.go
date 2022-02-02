package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/olumidesan/bookings/pkg/config"
	"github.com/olumidesan/bookings/pkg/handlers"
	"github.com/olumidesan/bookings/pkg/render"
)

// Create variable for app config
var app config.AppConfig

var session *scs.SessionManager

const portNumber = ":3000"

// main is the entry to the weba app
func main() {

	app.InProduction = false

	// 24 hour session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	// Store session in config
	app.Session = session

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
