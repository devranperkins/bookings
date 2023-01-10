package main

import (
	"github/dperkins/bookings/pkg/config"
	"github/dperkins/bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// Using Pat as a routing package
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux

	//using chi as a routing package

	mux := chi.NewRouter()
	mux.Use(NoSurf)

	//with chi we can use middleware

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
