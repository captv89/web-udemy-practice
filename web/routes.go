package main

import (
	"net/http"

	"github.com/captv89/web-udemy-practice/pkg/config"
	"github.com/captv89/web-udemy-practice/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	// Routes
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/contact", handler.Repo.Contact)
	return mux
}
