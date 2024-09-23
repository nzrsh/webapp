package main

import "github.com/julienschmidt/httprouter"

func SetupRouter() *httprouter.Router {
	r := httprouter.New()
	//HTML HANDLERS
	r.GET("/", homePageHandler)
	r.GET("/login", loginPageHandler)
	r.GET("/register", registerPageHandler)

	//API HANDLERS
	//r.POST("/auth", loginHandler)
	return r
}
