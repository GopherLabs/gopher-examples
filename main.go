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
		gopher.RENDERER: {
			"ViewsDir": "templates",
		},
	}
	app = gopher.NewApp(config)
	log = app.NewLog()
)

func main() {
	router := app.NewRouter()

	router.Get("/router", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Router!")
	})

	router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})
	router.Get("/handler", MyHandler)
	router.Post("/handler", MyHandler)
	router.Match("/verbs", MyHandler, "GET", "POST", "DELETE")
	router.All("/all", MyHandler)
	router.Get("/variables/{key}", PathParamHandler)
	router.Get("/view", ViewHandler)

	sample := app.NewSample()
	sample.SetName("Sample")
	app.NewLog().Info("sample is " + sample.GetName())

	subSample := sample.NewSample()
	subSample.SetName("SubSample")
	log.Info("subSample is " + subSample.GetName())

	subSubSample := sample.NewSample()
	subSubSample.SetName("subSubSample")
	log.Info("subSubSample is " + subSubSample.GetName())

	log.Info("sample is " + sample.GetName())

	router.Serve()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello Gophers from Handler! - HTTP Verb is: "+req.Method)
}

// Example of a handler that reads path parameters
func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "The Param Key is "+app.PathParam(req, "key"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	app.View(rw, http.StatusOK, "myview", nil)
}
