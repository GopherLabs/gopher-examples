package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

var (
	app = gopher.App()
	log = app.Log()
)

func main() {

	router := app.Router()
	addRoutes(router)
	router.Serve()
}

func addRoutes(router gopher.Routable) {

	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})

	router.Get("/handler", MyHandler)

	router.Get("/variables/{key}", PathParamHandler)

	router.Get("/view", ViewHandler)
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {

	log.Info("[%s] %s", req.Method, req.URL.Path)
	fmt.Fprint(rw, "Hello Gophers from Handler!")
}

//
// Example of a handler that reads path parameters
//
func PathParamHandler(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprint(rw, "Key is "+gopher.PathParam(req, "key"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {

	app.View(rw, http.StatusOK, "myview", nil)
}
