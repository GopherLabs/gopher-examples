package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	Route.Match("/hello", MatchHandler, []string{"GET", "POST", "PUT"})
	ListenAndServe()
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	Render.Text(w, "Hello, Y'all!")
}
