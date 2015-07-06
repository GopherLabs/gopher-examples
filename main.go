package main

import (
	"fmt"
	"net/http"

	"github.com/gopherlabs/gopher"
)

var (
	config = map[string]map[string]interface{}{
		gopher.LOGGER: {
			"FullTimestamp": false,
		},
		gopher.RENDERER: {
			"ViewsDir": "templates",
		},
	}
	app = gopher.NewApp(config)
	log = app.NewLog()
)

func main() {
	r := app.NewRouter()

	r.Get("/router", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Router!")
	})

	r.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Gophers!")
	})
	r.Get("/handler", MyHandler)
	r.Post("/handler", MyHandler)
	r.Match("/verbs", MyHandler, "GET", "POST", "DELETE")
	r.All("/all", MyHandler)
	r.Get("/variables/{key}", PathParamHandler)
	r.Get("/view", ViewHandler)

	//	sample := app.NewSample()
	//	sample.SetName("Sample")
	//	app.NewLog().Info("sample is " + sample.GetName())
	//
	//	subSample := sample.NewSample()
	//	subSample.SetName("SubSample")
	//	log.Info("subSample is " + subSample.GetName())
	//
	//	subSubSample := sample.NewSample()
	//	subSubSample.SetName("subSubSample")
	//	log.Info("subSubSample is " + subSubSample.GetName())
	//
	//	log.Info("sample is " + sample.GetName())

	r.Serve()
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
