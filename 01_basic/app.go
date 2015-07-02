package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

var (
	config = map[string]map[string]interface{}{
		gopher.LOGGER: {
			"FullTimestamp": true,
		},
	}
	app = gopher.App(config)
)

func main() {
	router := app.Router()
	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})
	router.Get("/handler", MyHandler)
	router.Get("/variables/{key}", PathParamHandler)
	router.Get("/view", ViewHandler)
	router.Serve()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	app.Log().Info("[%s] %s", req.Method, req.URL.Path)
	fmt.Fprint(rw, "Hello Gophers from Handler!")
}

// Example of a handler that reads path parameters
func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Key is "+app.PathParam(req, "key"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	app.View(rw, http.StatusOK, "myview", nil)
}
