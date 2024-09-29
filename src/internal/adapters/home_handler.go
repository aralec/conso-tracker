package adapters

import (
	"conso-tracker/src/external/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// RenderHome is a handler for the home page template
func (h *HomeHandler) RenderHome(w http.ResponseWriter, r *http.Request) {
	home := views.Home()
	home.Render(r.Context(), w)
}
