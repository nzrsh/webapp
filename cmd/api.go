package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

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
	//fmt.Fprintf(w, "Продукт с ID %d успешно обновлен.", id)
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

	imagePath := filepath.Join("public", "img", strconv.Itoa(id)+".jpg")
	err = os.Remove(imagePath)
	if err != nil {
		log.Printf("deleteProductHandler | Ошибка удаления изображения: %s\n", err)
	}

	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Продукт с ID %d успешно удален.", id)
}

func createProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Устанавливаем максимальный размер файла
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}

	// Получаем данные из формы
	typeProduct := r.FormValue("type")
	nameProduct := r.FormValue("name")
	priceProduct := r.FormValue("price")

	// Проверка на пустые значения
	if typeProduct == "" || nameProduct == "" || priceProduct == "" {
		http.Error(w, "Пожалуйста, заполните все поля", http.StatusBadRequest)
		return
	}

	// Обработка цены
	price, err := strconv.ParseFloat(priceProduct, 64)
	if err != nil {
		http.Error(w, "Некорректная цена", http.StatusBadRequest)
		return
	}

	// Сохранение картинки
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Ошибка при получении изображения", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создаем продукт
	product := Product{
		ID:    0,
		Type:  typeProduct,
		Name:  nameProduct,
		Price: price,
	}

	// Создание ID для продукта (можно использовать автоинкремент из базы данных)
	productID, err := CreateProductInTable(product)
	if err != nil {
		log.Printf("createProductHandler | Ошибка при создании продукта: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Путь для сохранения изображения
	imgPath := filepath.Join("public", "img", fmt.Sprintf("%d.jpg", productID))

	// Создаем файл для сохранения изображения
	out, err := os.Create(imgPath)
	if err != nil {
		http.Error(w, "Ошибка при сохранении изображения", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Копируем содержимое файла в новый файл
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Ошибка при копировании изображения", http.StatusInternalServerError)
		return
	}
	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product) // Возвращаем созданный продукт
}

func getProductImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !strings.HasSuffix(id, ".jpg") {
		id += ".jpg" // Добавляем .jpg, если его нет
	}

	imagePath := filepath.Join("public", "img", id)
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, imagePath)

	if err := r.Context().Err(); err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
}
