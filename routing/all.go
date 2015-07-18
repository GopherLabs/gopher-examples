package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.All("/", AllHandler)
	ListenAndServe()
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	Render.Text(w, "Hello, "+r.Method)
}
