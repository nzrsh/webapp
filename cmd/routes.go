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
	r.GET("/storage", JWTAuthMiddleware(storagePageHandler))
	r.GET("/add", JWTAuthMiddleware(addPageHandler))

	//AUTH HANDLERS

	r.POST("/auth/register", registerHandler)
	r.POST("/auth/login", loginHandler)
	r.POST("/auth/logout", logoutHandler)

	//Выдача данных о пользователе по токену
	r.GET("/auth/me", meHandler)

	//API PRODUCT HANDLERS
	//Эндпоинт получения списка продуктов в формате массива объектов json
	r.GET("/api/products", getProductsHandler)
	//Эндпоинт получения продукта в формате объекта json
	r.GET("/api/products/:id", getProductHandler)
	//Эндпоинт обновления продукта
	r.POST("/api/products/:id", JWTAuthMiddleware(updateProductHandler))
	//Эндпоинт удаления продукта
	r.DELETE("/api/products/:id", JWTAuthMiddleware(deleteProductHandler))
	//Эндпоинт добавления продукта
	r.POST("/api/products", JWTAuthMiddleware(createProductHandler))
	//Картинка продукта
	r.GET("/api/img/:id", getProductImageHandler)

	//API STORAGE HANDLERS
	r.GET("/storage/files", JWTAuthMiddleware(getFilesHandler))
	r.POST("/storage/files/upload", JWTAuthMiddleware(uploadFileHandler))
	r.GET("/storage/files/image", JWTAuthMiddleware(getImageHandler))
	r.DELETE("/storage/files/:filename", JWTAuthMiddleware(deleteFileHandler))
	r.PUT("/storage/files/:filename", JWTAuthMiddleware(renameFileHandler))
	return r
}
