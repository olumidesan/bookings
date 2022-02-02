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
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
