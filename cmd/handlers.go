package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/home.html")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/login.html")
}

func registerPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/reg.html")
}
