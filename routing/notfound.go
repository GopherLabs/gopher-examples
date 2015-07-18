package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, Gopher!")
	})
	Route.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Could not find page")
	})
	ListenAndServe()
}
