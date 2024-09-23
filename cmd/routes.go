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

	//Эндпоинт получения списка продуктов в формате массива объектов json
	r.GET("/api/products", getProducts)
	//Эндпоинт получения продукта в формате объекта json
	r.GET("/api/products/:id", getProduct)

	return r
}
