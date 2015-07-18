package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Context.Set("user", "Ricardo Rossi")
	Route.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, "+Context.Get("user").(string))
	})
	ListenAndServe()
}
