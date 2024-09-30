package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func getFilesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Получение логина пользователя из токена
	username, err := getLoginFromCookie(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Путь к директории с файлами
	uploadDir := filepath.Join("uploads", username)

	// Проверка, существует ли директория
	_, err = os.Stat(uploadDir)
	if os.IsNotExist(err) {
		// Директория не существует, возвращаем пустой массив
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]")) // Возвращаем пустой JSON массив
		return
	} else if err != nil {
		// Если произошла другая ошибка при проверке директории
		http.Error(w, "Error accessing directory", http.StatusInternalServerError)
		return
	}

	// Получаем список файлов
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}

	if len(files) == 0 {
		// Если файлов нет, возвращаем пустой массив
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]")) // Возвращаем пустой JSON массив
		return
	}

	var fileInfos []FileInfo
	for _, file := range files {
		if !file.IsDir() {
			fileInfo, err := os.Stat(filepath.Join(uploadDir, file.Name()))
			if err != nil {
				continue
			}
			fileInfos = append(fileInfos, FileInfo{
				Name:    fileInfo.Name(),
				Size:    fileInfo.Size(),
				ModTime: fileInfo.ModTime().Format(time.RFC3339),
				IsImage: isImage(fileInfo.Name()),
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileInfos) // Отправляем информацию о файлах
}

// Функция для проверки, является ли файл изображением
func isImage(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Проверка, что метод запроса - POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получение логина пользователя из токена
	username, err := getLoginFromCookie(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Путь к директории с файлами
	uploadDir := filepath.Join("uploads", username)

	// Создание директории, если она не существует
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	// Чтение файла из запроса
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создание файла на сервере с использованием имени загруженного файла
	dst, err := os.Create(filepath.Join(uploadDir, fileHeader.Filename)) // Используем имя загруженного файла
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копирование содержимого загруженного файла в файл на сервере
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Успешный ответ
	w.WriteHeader(http.StatusCreated) // Возвращаем статус 201 Created
	w.Write([]byte("File uploaded successfully"))
}

func getImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Получение логина пользователя из токена
	username, err := getLoginFromCookie(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Путь к файлу
	fileName := r.URL.Query().Get("filename")
	filePath := filepath.Join("uploads", username, fileName)

	// Проверка существования файла
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Установка заголовков
	w.Header().Set("Content-Type", "image/jpeg") // или другой тип изображения
	http.ServeFile(w, r, filePath)               // Отправляем файл клиенту
}

func renameFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, err := getLoginFromCookie(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	oldName := ps.ByName("filename")
	var requestBody struct {
		NewName string `json:"newName"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.NewName == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	oldFilePath := filepath.Join("uploads", username, oldName)
	newFilePath := filepath.Join("uploads", username, requestBody.NewName)

	err = os.Rename(oldFilePath, newFilePath)
	if err != nil {
		http.Error(w, "Error renaming file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Успешное переименование файла
}

func deleteFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, err := getLoginFromCookie(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	filename := ps.ByName("filename")
	filePath := filepath.Join("uploads", username, filename)

	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Успешное удаление файла
}
