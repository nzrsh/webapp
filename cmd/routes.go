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
	r.GET("/api/products", getProductsHandler)
	//Эндпоинт получения продукта в формате объекта json
	r.GET("/api/products/:id", getProductHandler)
	//Эндпоинт обновления продукта
	r.PUT("/api/products/:id", updateProductHandler)
	//Эндпоинт удаления продукта
	r.DELETE("/api/products/:id", deleteProductHandler)
	//Эндпоинт добавления продукта
	r.POST("/api/products", createProductHandler)

	return r
}
