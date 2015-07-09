package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

func main() {
	app := gopher.NewApp()
	router := app.NewRouter()
	router.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "Hello, Gopher!")
	})
	router.Serve()
}
