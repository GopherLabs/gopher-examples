package main

import "github.com/gopherlabs/gopher"

type WebContext struct{}

func main() {
	gopher.Start()
	gopher.Hello("App")
	router := gopher.NewRouter(WebContext{})
	_ = router
	//router.Get("/", (*WebContext).Home)
}

//func (c *WebContext) Home(rw web.ResponseWriter, req *web.Request) {
//	if c.User != nil {
//		fmt.Fprint(rw, "Hello,", c.User.Name)
//	} else {
//		fmt.Fprint(rw, "Hello, anonymous person")
//	}
//}
