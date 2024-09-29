package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// PAGES

func homePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/home.html")
}

func registerPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("token")
	if err == nil && cookie != nil {
		tokenString := cookie.Value

		// Проверяем токен
		if _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("svo"), nil
		}); err == nil {
			// Токен действителен, перенаправляем на главную страницу
			http.Redirect(w, r, "/?message=authorized", http.StatusFound)
			return
		}
	}
	http.ServeFile(w, r, "./public/html/reg.html")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	cookie, err := r.Cookie("token")
	if err == nil && cookie != nil {
		tokenString := cookie.Value

		// Проверяем токен
		if _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("svo"), nil
		}); err == nil {
			// Токен действителен, перенаправляем на главную страницу
			http.Redirect(w, r, "/?message=authorized", http.StatusFound)
			return
		}
	}

	http.ServeFile(w, r, "./public/html/login.html")
}

func storagePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/storage.html")
}
