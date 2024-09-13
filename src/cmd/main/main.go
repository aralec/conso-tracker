package main

import (
	"conso-tracker/src/cmd/render"
	"net/http"
)

// main runs the core application (server-side API)
func main() {
	r := render.GetRenderingMux() // TODO : Main should start only the back-end API
	http.ListenAndServe(":8080", r)
}
