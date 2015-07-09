package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

func main() {
	r := gopher.NewApp().NewRouter()
	r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "Hello, Gopher!")
	})
	r.Serve()
}
