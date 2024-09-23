package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/crypto/bcrypt"
)

func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Подключение в базе данных установлено, путь: ", path)
	return db
}

func GetProductsFromTable() []Product {
	var Products []Product

	if DB == nil {
		log.Fatal("Соединение с базой данных не было установлено.")
	}

	rows, err := DB.Query("SELECT id, type, name, price FROM products")
	if err != nil {
		log.Fatalf("Ошибка при получении списка продуктов: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Type, &p.Name, &p.Price)
		if err != nil {
			log.Fatalf("Ошибка при считывании продукта в структуру: %s", err)
		}
		Products = append(Products, p)

		if err := rows.Err(); err != nil {
			log.Fatal("Ошибка при выполнении запроса на чтение продуктов: ", err)
		}

	}
	return Products
}
