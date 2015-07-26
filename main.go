package main

import (
	"net/http"
	"time"

	. "github.com/gopherlabs/gopher"
)

type MyContext struct {
	Username string
}

func main() {

	App.Config(Config{
		KEY_ROUTER: ConfigRouter{
			Port: 3002,
			Host: "0.0.0.0",
		},
		KEY_LOGGER: ConfigLogger{
			TimestampFormat: time.RFC822,
			LogLevel:        LEVEL_DEBUG,
		},
	})

	App.Use(MyAppMiddleWareFunc1)

	Route.Use(MyMiddleWareFunc1, MyContext{Username: "Ricardo"})
	Route.Use(MyMiddleWareFunc2)

	Route.Get("/route", func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Hello Route!")
	})

	Route.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Hello, Gophers!")
	})
	Route.Get("/handler", MyHandler)
	Route.Post("/handler", MyHandler, MyAppMiddleWareRouteHanlder)
	Route.Match("/verbs", MyHandler, []string{"GET", "POST", "DELETE"}, MyAppMiddleWareRouteHanlder)
	Route.All("/all", MyHandler)
	Route.Get("/variables/{key}", PathParamHandler)
	Route.Get("/view", ViewHandler)
	Route.Get("/route", MyHandler, MyAppMiddleWareRouteHanlder, MyAppMiddleWareRouteHanlder2)

	Route.NotFound(func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Could not find page")
	})

	group := RouteGroup.New(GroupMatcher{
		PathPrefix: "/abc",
	})
	group.Get("/group", func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Hello Group!")
	})
	group.Use(MyMiddleWareFunc1, MyContext{Username: "Ricardo"})

	adminGroup := RouteGroup.New(GroupMatcher{
		PathPrefix: "/admin",
		Queries:    []string{"q", "boo"},
	})
	adminGroup.Get("/profile", func(rw http.ResponseWriter, req *http.Request) {
		Render.Text(rw, "Hello Admin!")
	})

	ListenAndServe()
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	Render.Text(rw, "Hello Gophers from Handler! - HTTP Verb is: "+req.Method)
}

// Example of a handler that reads path parameters
func PathParamHandler(w http.ResponseWriter, r *http.Request) {
	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	Log.Info("The Key : " + Route.Var(r, "key"))
	Render.Text(w, "The Key : "+Route.Var(r, "key"))
	user := Context.Get(r, "user").(*MyContext)
	Log.Info("Inside PathParamHandler = My username is %s ", user.Username)

	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	Log.Info("Removing key... ")
	Context.Delete(r, "user")
	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	Log.Info("Cool, I am logging!")
}

func ViewHandler(rw http.ResponseWriter, req *http.Request) {
	Render.View(rw, "myview", nil)
}

func MyAppMiddleWareFunc1(rw http.ResponseWriter, r *http.Request, next func(), args ...interface{}) {
	Log.Info("======== Inside My APP MyMiddleWareFunc 1")
	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	user := new(MyContext)
	user.Username = "rrossi"
	Context.Set(r, "user", user)
	next()
}

func MyMiddleWareFunc1(rw http.ResponseWriter, r *http.Request, next func(), args ...interface{}) {
	Log.Info("Inside My MyMiddleWareFunc 1")
	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	if len(args) > 0 {
		Log.Info("Hello %s ", args[0].(MyContext).Username)
	}
	if Context.Has(r, "user") {
		user := Context.Get(r, "user").(*MyContext)
		Log.Info("My username is %s ", user.Username)
		user.Username = "Modified " + user.Username
	}
	next()
}

func MyMiddleWareFunc2(rw http.ResponseWriter, r *http.Request, next func(), args ...interface{}) {
	Log.Info("Inside My MyMiddleWareFunc 2")
	Log.Info("Has user key? %t ", Context.Has(r, "user"))
	if Context.Has(r, "user") {
		user := Context.Get(r, "user").(*MyContext)
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
