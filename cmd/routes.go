package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter() *httprouter.Router {
	r := httprouter.New()

	r.ServeFiles("/public/*filepath", http.Dir("public"))

	//HTML HANDLERS
	r.GET("/", homePageHandler)
	r.GET("/login", loginPageHandler)
	r.GET("/register", registerPageHandler)

	//API HANDLERS
	//r.POST("/auth", loginHandler)
	return r
}
