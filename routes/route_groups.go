package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {

	Context.Set("user", "Ricardo Rossi")

	Route.Get("/text", func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Hello Text")
	})

	Route.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		Log.Info("Now we are cooking!")
		Render.Text(w, "Hello, "+Context.Get("user").(string))
	})

	Route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		Render.Text(w, "Could not find this page")
	})

	ListenAndServe()
}
