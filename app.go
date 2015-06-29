package main

import (
	"fmt"
	"github.com/gopherlabs/gopher"
	"net/http"
)

func main() {
	gopher.Start()
	gopher.Hello("App")

	router := gopher.NewRouter()

	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})

	router.Get("/handler", HomeHandler)

	router.Serve()
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello Gophers in Handler!")
}
