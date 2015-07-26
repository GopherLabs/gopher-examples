package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		Context.Set(r, "name", "Ricardo Rossi")
		Render.Text(w, "Hello, "+Context.Get(r, "name").(string))
	})
	ListenAndServe()
}
