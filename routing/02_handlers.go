package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Get("/hello", HelloHandler)
	ListenAndServe()
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	Render.Text(w, "Hello, Handler!")
}
