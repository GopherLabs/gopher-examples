package main

import (
	"fmt"
	"net/http"

	. "github.com/gopherlabs/gopher"
)

type MyContext struct {
	Username string
}

func main() {

	var config = map[string]map[string]interface{}{
		App.LOGGER: {
			"FullTimestamp": false,
		},
		App.RENDERER: {
			"ViewsDir": "templates",
		},
	}
	_ = config

	App.Config(config)
	App.Use(MyAppMiddleWareFunc1)

	Route.Use(MyMiddleWareFunc1, MyContext{Username: "Ricardo"})
	Route.Use(MyMiddleWareFunc2)

	Route.Get("/Route", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Route!")
	})

	Route.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello, Gophers!")
	})
	Route.Get("/handler", MyHandler)
	Route.Post("/handler", MyHandler, MyAppMiddleWareRouteHanlder)
	Route.Match("/verbs", MyHandler, []string{"GET", "POST", "DELETE"}, MyAppMiddleWareRouteHanlder)
	Route.All("/all", MyHandler)
	Route.Get("/variables/{key}", PathParamHandler)
	Route.Get("/view", ViewHandler)
	Route.Get("/route", MyHandler, MyAppMiddleWareRouteHanlder, MyAppMiddleWareRouteHanlder2)

	Route.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Could not find page")
	})

	/*
		sub := r.SubRouter()
		sub.Get("/shirts", func(rw http.ResponseWriter, req *http.Request) {
			fmt.Fprint(rw, "sub Shirt")
		})

		subSub := sub.SubRouter()
		subSub.Get("/shirts", func(rw http.ResponseWriter, req *http.Request) {
			fmt.Fprint(rw, "subSub Shirt")
		})
	*/

	ListenAndServe()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	Render.Text(rw, "Hello Gophers from Handler! - HTTP Verb is: "+req.Method)
}

// Example of a handler that reads path parameters
func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	Log.Info("Has user key? %t ", Context.Has("user"))
	Log.Info("The Key : " + Route.Var(req, "key"))
	Render.Text(rw, "The Key : "+Route.Var(req, "key"))
	user := Context.Get("user").(*MyContext)
	Log.Info("Inside PathParamHandler = My username is %s ", user.Username)

	Log.Info("Has user key? %t ", Context.Has("user"))
	Log.Info("Removing key... ")
	Context.Remove("user")
	Log.Info("Has user key? %t ", Context.Has("user"))
	Log.Info("Cool, I am logging!")
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	Render.View(rw, "myview", nil)
}

func MyAppMiddleWareFunc1(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	Log.Info("======== Inside My APP MyMiddleWareFunc 1")
	Log.Info("Has user key? %t ", Context.Has("user"))
	user := new(MyContext)
	user.Username = "rrossi"
	Context.Set("user", user)
	next()
}

func MyMiddleWareFunc1(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	Log.Info("Inside My MyMiddleWareFunc 1")
	Log.Info("Has user key? %t ", Context.Has("user"))
	if len(args) > 0 {
		Log.Info("Hello %s ", args[0].(MyContext).Username)
	}
	if Context.Has("user") {
		user := Context.Get("user").(*MyContext)
		Log.Info("My username is %s ", user.Username)
		user.Username = "Modified " + user.Username
	}
	next()
}

func MyMiddleWareFunc2(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	Log.Info("Inside My MyMiddleWareFunc 2")
	Log.Info("Has user key? %t ", Context.Has("user"))
	if Context.Has("user") {
		user := Context.Get("user").(*MyContext)
		user.Username = "modified again"
		Log.Info("My username is %s ", user.Username)
	}
	next()
}

func MyAppMiddleWareRouteHanlder(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	Log.Info("From MyAppMiddleWareRouteHanlder")
	next()
}

func MyAppMiddleWareRouteHanlder2(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	Log.Info("From MyAppMiddleWareRouteHanlder2")
	next()
}
