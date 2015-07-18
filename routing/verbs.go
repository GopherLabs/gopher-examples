package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/products", VerbHandler)
	Route.Post("/form", VerbHandler)
	Route.Put("/update", VerbHandler)
	Route.Delete("/etc", VerbHandler)
	Route.Head("/etc", VerbHandler)
	Route.Options("/etc", VerbHandler)
	ListenAndServe()
}

func VerbHandler(w http.ResponseWriter, r *http.Request) {
	Render.Text(w, "Hello, "+r.Method)
}
