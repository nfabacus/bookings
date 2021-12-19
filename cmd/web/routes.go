package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nfabacus/bookings/internal/config"
	"github.com/nfabacus/bookings/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/", handlers.Repo.PostForm)
	mux.Get("/submission-summary", handlers.Repo.SubmissionSummary)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/example-json", handlers.Repo.GetExampleJSON)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
