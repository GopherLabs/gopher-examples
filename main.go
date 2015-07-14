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
		"LOGGER": {
			"FullTimestamp": false,
		},
		"RENDERER": {
			"ViewsDir": "templates",
		},
	}
	_ = config

	// TODO Replace with App.Config()
	Config(config)
	App.Use(MyAppMiddleWareFunc1)

	Router.Use(MyMiddleWareFunc1, MyContext{Username: "Ricardo"})
	Router.Use(MyMiddleWareFunc2)

	Router.Get("/router", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello Router!")
	})

	Router.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "Hello, Gophers!")
	})
	Router.Get("/handler", MyHandler)
	Router.Post("/handler", MyHandler, MyAppMiddleWareRouteHanlder)
	Router.Match("/verbs", MyHandler, []string{"GET", "POST", "DELETE"}, MyAppMiddleWareRouteHanlder)
	Router.All("/all", MyHandler)
	Router.Get("/variables/{key}", PathParamHandler)
	Router.Get("/view", ViewHandler)
	Router.Get("/route", MyHandler, MyAppMiddleWareRouteHanlder, MyAppMiddleWareRouteHanlder2)

	Router.NotFound(func(rw http.ResponseWriter, req *http.Request) {
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

	Log.Info("Serve(), I am logging!")
	ListenAndServe()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	Render.Text(rw, "Hello Gophers from Handler! - HTTP Verb is: "+req.Method)
}

// Example of a handler that reads path parameters
func PathParamHandler(rw http.ResponseWriter, req *http.Request) {
	Log.Info("Has user key? %t ", Context.Has("user"))
	//Log.Info("The Param Key is "+.PathParam(req, "key"))
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
