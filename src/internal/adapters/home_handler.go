package adapters

import (
	"conso-tracker/src/external/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	home := views.Home()
	home.Render(r.Context(), w)
}
