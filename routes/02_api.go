package main

import (
	"net/http"

	g "github.com/gopherlabs/gopher"
)

func main() {
	g.Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		g.Render.Text(w, "Hello, Gopher!")
	})
	g.ListenAndServe()
}
