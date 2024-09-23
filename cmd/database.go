package main

import (
	"database/sql"
	"fmt"
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

	_, err = db.Exec(CreateTableUsersQuery)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы пользователей: %s", err)
	}

	_, err = db.Exec(CreateTableProductsQuery)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы продуктов: %s", err)
	}
	return db
}

// Функция для добавления пользователя в таблицу
func LoadDefaultData() {

	products := []Product{
		{1, "motherboard", "GIGABYTE h610m v2 DDR4", 8900.0},
		{2, "gpu", "MSI 1050 TI 8GB", 14990.0},
		{3, "cpu", "Xeon E3-1230 v2", 1500.0},
		{4, "motherboard", "ASUS ROG Strix B450-F Gaming", 10900.0},
		{5, "gpu", "AMD Radeon RX 580 8GB", 17990.0},
		{6, "cpu", "AMD Ryzen 5 3600", 9500.0},
		{7, "motherboard", "MSI B450M PRO-VDH MAX", 6700.0},
		{8, "gpu", "NVIDIA GeForce GTX 1660 Super 6GB", 19990.0},
		{9, "cpu", "Intel Core i5-10400", 8500.0},
	}

	for _, product := range products {
		stmt, err := DB.Prepare("INSERT OR IGNORE INTO products(id, type, name, price) VALUES(?, ?, ?, ?)")
		if err != nil {
			log.Fatalf("Ошибка подготовки стандартных значений для базы данных: %s", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.ID, product.Type, product.Name, product.Price)
		if err != nil {
			log.Fatalf("Ошибка добавления стандартных значений в таблицу продуктов: %s", err)
		}

	}

	log.Printf("Начальные значения для продуктов добавлены или уже существуют.")
}

// Получение списка продуктов из базы
func GetProductsFromTable() ([]Product, error) {
	var Products []Product
	rows, err := DB.Query("SELECT id, type, name, price FROM products")
	if err != nil {
		fmt.Errorf("Ошибка при получении списка продуктов: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Type, &p.Name, &p.Price)
		if err != nil {
			fmt.Errorf("ошибка при считывании продукта в структуру: %s", err)
		}
		Products = append(Products, p)

		if err := rows.Err(); err != nil {
			fmt.Errorf("ошибка при выполнении запроса на чтение продуктов: %s", err)
		}

	}
	return Products, nil
}

// Получение продукта по его ID
func GetProductFromTable(id int) (Product, error) {
	var p Product

	err := DB.QueryRow("SELECT id, type, name, price FROM products WHERE id = ?", id).Scan(&p.ID, &p.Type, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("продукт с ID %d не найден", id)
		}
		return p, err
	}
	return p, nil
}
