package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/dgrijalva/jwt-go"
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

// API
// Отдаёт пользователю JSON массив объектов Product, исключая изображение
func getProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products, err := GetProductsFromTable()
	if err != nil {
		log.Printf("getProductsHandler | Ошибка получения списка продуктов: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Printf("getProductsHandler | Ошибка при кодировании списка продуктов: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Отдаёт пользователю JSON объект продукта по его ID
func getProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("getProductHandler | Ошибка при парсинге айди: %s\n", err)
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	product, err := GetProductFromTable(id)
	if err != nil {
		log.Printf("getProductHandler | Ошибка при получении продукта: %s\n", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

func updateProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("updateProductHandler | Ошибка при получении ID продукта: %s\n", err)
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Printf("updateProductHandler | Ошибка при десереализации продукта: %s\n", err)
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	// Обновляем продукт в базе данных
	err = UpdateProductFromTable(id, product)
	if err != nil {
		log.Printf("updateProductHandler | Ошибка при обновлении продукта: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Продукт с ID %d успешно обновлен.", id)
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("deleteProductHandler | Ошибка при получении ID продукта: %s\n", err)
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}
	err = DeleteProductFromTable(id)
	if err != nil {
		log.Printf("deleteProductHandler | Ошибка удаления продукта: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Продукт с ID %d успешно удален.", id)
}

func createProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Printf("createProductHandler | Ошибка при получении ID продукта: %s\n", err)
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	id, err := CreateProductInTable(product)
	if err != nil {
		log.Printf("createProductHandler | Ошибка создании продукта: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Продукт %d %s успешно создан.", id, product.Name)
}
