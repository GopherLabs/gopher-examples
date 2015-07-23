package main

import (
	"net/http"

	. "github.com/gopherlabs/gopher"
)

func main() {
	group := RouteGroup.New(GroupMatcher{
		PathPrefix: "/products",
	})
	group.Get("/group", func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Hello Group!")
	})
	ListenAndServe()
}
