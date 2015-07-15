package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, Gopher!")
	})
	ListenAndServe()
}
