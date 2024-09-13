package main

import (
	"conso-tracker/src/cmd/render"
	"net/http"
)

// main runs the core application (server-side code + templating)
func main() {
	r := render.GetRenderingMux()
	http.ListenAndServe(":8080", r)
}
