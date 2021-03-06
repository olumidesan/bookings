package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/olumidesan/bookings/pkg/config"
	"github.com/olumidesan/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Using chi router
	mux := chi.NewRouter()

	// Catch panics and stack-trace
	mux.Use(middleware.Recoverer)

	// Protect against CSRF
	mux.Use(NoSurf)

	// Load session data
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Serve files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
