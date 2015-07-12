package main

import (
	"fmt"
	"net/http"

	g "github.com/gopherlabs/gopher"
)

func main() {

	g.Initialize()

	g.Context.Set("user", "Ricardo Rossi")

	g.Router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		g.Log.Info("Now we are cooking!")
		fmt.Fprintln(rw, "Hello, "+g.Context.Get("user").(string))
	})

	g.Router.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "hey! Could not find this page")
	})

	//app.Routes().Get
	//	router.Get("/", func(rw http.ResponseWriter, req *http.Request) {
	//		fmt.Fprintln(rw, "Hello, Gopher!")
	//	})
	//	router.Serve()

	g.ListenAndServe()
}
