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

type MyMiddleware struct {
}

func MyAppMiddleWareFunc1(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	fmt.Fprint(rw, "\n\nInside My APP MyMiddleWareFunc 1 \n")
	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
	user := new(MyContext)
	user.Username = "rrossi"
	app.Context().Write("user", user)
	next()
}

func MyMiddleWareFunc1(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	fmt.Fprint(rw, "\n\nInside My MyMiddleWareFunc 1\n")
	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
	if len(args) > 0 {
		fmt.Fprintf(rw, "Hello %s \n", args[0].(MyContext).Username)
	}
	user := app.Context().Read("user").(*MyContext)
	fmt.Fprintf(rw, "My username is %s \n", user.Username)
	user.Username = "Modified " + user.Username
	next()
}

func MyMiddleWareFunc2(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	fmt.Fprint(rw, "\n\nInside My MyMiddleWareFunc 2 \n")
	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
	user := app.Context().Read("user").(*MyContext)
	user.Username = "modified again"
	fmt.Fprintf(rw, "My username is %s \n\n", user.Username)
	next()
}

func MyAppMiddleWareRouteHanlder(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	fmt.Fprint(rw, "\n\n From MyAppMiddleWareRouteHanlder \n")
	next()
}

func MyAppMiddleWareRouteHanlder2(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	fmt.Fprint(rw, "\n\n From MyAppMiddleWareRouteHanlder2 \n")
	next()
}

type MyContext struct {
	Username string
}

func main() {

	app.Use(MyAppMiddleWareFunc1)

	r := app.NewRouter()

	r.Use(MyMiddleWareFunc1, MyContext{Username: "Ricardo"})
	r.Use(MyMiddleWareFunc2)

	r.Get("/router", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Router!")
	})

	r.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello, Gophers!")
	})
	r.Get("/handler", MyHandler)
	r.Post("/handler", MyHandler, MyAppMiddleWareRouteHanlder)
	//r.Match("/verbs", MyHandler, "GET", "POST", "DELETE")
	r.All("/all", MyHandler)
	r.Get("/variables/{key}", PathParamHandler)
	r.Get("/view", ViewHandler)
	r.Get("/route", MyHandler, MyAppMiddleWareRouteHanlder, MyAppMiddleWareRouteHanlder2)

	r.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Could not find page")
	})

	sub := r.SubRouter()
	sub.Get("/shirts", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "sub Shirt")
	})

	subSub := sub.SubRouter()
	subSub.Get("/shirts", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "subSub Shirt")
	})

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
	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
	fmt.Fprint(rw, "The Param Key is "+app.PathParam(req, "key"))
	user := app.Context().Read("user").(*MyContext)
	fmt.Fprintf(rw, "\nInside PathParamHandler = My username is %s \n", user.Username)

	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
	fmt.Fprint(rw, "Removing key... \n")
	app.Context().Remove("user")
	fmt.Fprintf(rw, "Has user key? %t \n", app.Context().Has("user"))
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	app.View(rw, http.StatusOK, "myview", nil)
}
