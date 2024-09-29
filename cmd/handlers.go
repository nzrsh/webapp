package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PAGES

func homePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/home.html")
}

func registerPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/reg.html")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/login.html")
}

func storagePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/storage.html")
}

func editPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/edit.html")
}

func addPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/add.html")
}
