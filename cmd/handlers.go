package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// PAGES

func homePageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/home.html")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/login.html")
}

func registerPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/html/reg.html")
}

// API
// Отдаёт пользователю JSON массив объектов Product, исключая изображение
func getProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products, err := GetProductsFromTable()
	if err != nil {
		log.Printf("Ошибка getProducts: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Printf("Ошибка при кодировании списка продуктов: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	product, err := GetProductFromTable(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(product)

}
