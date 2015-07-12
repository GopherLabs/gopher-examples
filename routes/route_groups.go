package main

import (
	"fmt"
	"net/http"

	g "github.com/gopherlabs/gopher"
)

func main() {

	g.Initialize()

	g.Log.Debug("hello debug() we are cooking!")
	g.Log.Info("Now we are cooking!")
	g.Log.Warn("Now we are cooking!")

	g.Router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "Hello, hello!")
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
