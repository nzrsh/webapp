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

	//AUTH HANDLERS

	r.POST("/auth/register", registerHandler)
	r.POST("/auth/login", loginHandler)

	//API HANDLERS

	//Эндпоинт получения списка продуктов в формате массива объектов json
	r.GET("/api/products", JWTAuthMiddleware(getProductsHandler))
	//Эндпоинт получения продукта в формате объекта json
	r.GET("/api/products/:id", JWTAuthMiddleware(getProductHandler))
	//Эндпоинт обновления продукта
	r.PUT("/api/products/:id", JWTAuthMiddleware(updateProductHandler))
	//Эндпоинт удаления продукта
	r.DELETE("/api/products/:id", JWTAuthMiddleware(deleteProductHandler))
	//Эндпоинт добавления продукта
	r.POST("/api/products", JWTAuthMiddleware(createProductHandler))

	return r
}
