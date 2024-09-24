package render

import (
	"conso-tracker/src/internal/adapters"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// GetRenderingMux creates a new router for the application that serves views rendered by Go Templ
func GetRenderingMux() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	homeHandler := adapters.NewHomeHandler()
	r.Get("/home", homeHandler.Home)

	fileHandler := adapters.NewFileHandler()
	r.Get("/importer", fileHandler.Importer)

	return r
}
