package main

import (
	"net/http"
	// Only for the purpose of syntax clarity, we are using dot import notation below.
	// It should be noted, however, that the Go team does not recommend using the dot
	// import since it can cause some odd behaviour in certain cases.
	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello, Gopher!")
	})
	ListenAndServe()
}
