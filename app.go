package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

func main() {
	gopher.Start()
	router := gopher.NewRouter()

	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})

	router.Get("/handler", MyHandler)

	router.Get("/variables/{key}", PathParamHandler)

	router.Get("/view", ViewHandler)

	router.Serve()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	gopher.NewLogger().Info("[%s] %s", req.Method, req.URL.Path)
	fmt.Fprint(rw, "Hello Gophers from Handler!")
}

func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Key is "+gopher.PathParam(req, "key"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	gopher.View(rw, http.StatusOK, "myview", nil)
}
