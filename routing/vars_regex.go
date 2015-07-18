package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/{var}", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, "+Route.Var(r, "var"))
	})
	ListenAndServe()
}
