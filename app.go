package main

import (
	"fmt"
	"github.com/gopherlabs/gopher"
	"net/http"
)

func main() {
	gopher.Start()
	router := gopher.NewRouter()

	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})

	router.Get("/handler", MyHandler)

	router.Get("/variables/{key}", PathParamHandler)

	router.Serve()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	gopher.NewLogger().Info("[%s] %s", req.Method, req.URL.Path)
	fmt.Fprint(rw, "Hello Gophers from Handler!")
}

func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Key is "+gopher.PathParam(req, "key"))
}
