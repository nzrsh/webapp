package main

import (
	"database/sql"
	"log"
	"net/http"
)

var DB *sql.DB

// Initial
func main() {
	log.Println("Попытка запуститься...")
	addr := "localhost:8080"
	databasePath := "./database/data.db"
	r := SetupRouter()
	DB = OpenDatabase(databasePath)
	LoadDefaultData()
	defer DB.Close()

	log.Println("Сервер запущен по адресу:", addr)
	log.Fatal(http.ListenAndServe(":8080", r))
}
