package main

import (
	"fmt"
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {

	Context.Set("user", "Ricardo Rossi")

	Router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		Log.Info("Now we are cooking!")
		fmt.Fprintln(rw, "Hello, "+Context.Get("user").(string))
	})

	Router.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "hey! Could not find this page")
	})

	ListenAndServe()
}
