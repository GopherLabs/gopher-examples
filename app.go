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
	router.Get("/", Home)
	router.Serve()
}

func Home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello, Gophers")
}
