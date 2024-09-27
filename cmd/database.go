package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

		_, err = stmt.Exec(product.ID, product.Type, product.Name, product.Price)
		stmt.Close()
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
		return nil, fmt.Errorf("GetProductsFromTable | ошибка при получении списка продуктов: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Type, &p.Name, &p.Price)
		if err != nil {
			return nil, fmt.Errorf("GetProductsFromTable |  ошибка при считывании продукта в структуру: %s", err)
		}
		Products = append(Products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetProductsFromTable | ошибка при выполнении запроса на чтение продуктов: %s", err)
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

// Изменение продукта в БД по его ID
func UpdateProductFromTable(id int, product Product) error {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id=?)", id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("UpdateProductFromTable | ошибка при проверке существования продукта: %s", err)
	}
	if !exists {
		return fmt.Errorf("продукт с ID %d не найден", id)
	}
	stmt, err := DB.Prepare("UPDATE products SET type = ?, name = ?, price = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("UpdateProductFromTable |  ошибка обновления продукта %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Type, product.Name, product.Price, id)
	if err != nil {
		return fmt.Errorf("UpdateProductFromTable | ошибка обновления продукта %s", err)
	}
	return nil
}

// Удаление продукта из БД по его ID
func DeleteProductFromTable(id int) error {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id=?)", id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("DeleteProductFromTable | ошибка при проверке существования продукта: %s", err)
	}
	if !exists {
		return fmt.Errorf("продукт с ID %d не найден", id)
	}

	stmt, err := DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return fmt.Errorf("DeleteProductFromTable | ошибка удаления продукта %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("DeleteProductFromTable | ошибка удаления продукта %s", err)
	}

	return nil
}

// Добавление нового продукта в качестве новой записи в БД
func CreateProductInTable(product Product) (int, error) {

	stmt, err := DB.Prepare("INSERT INTO products (type, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreateProductInTable | ошибка cоздания продукта %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Type, product.Name, product.Price)
	if err != nil {
		return 0, fmt.Errorf("CreateProductInTable | ошибка cоздания продукта %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateProductInTable | ошибка cоздания продукта %s", err)
	}

	return int(id), nil
}

// Хеширование пароля с помощью безопасного алгоритма шифрования bcrypt
func HashPassword(password string) (string, error) {
	//Генерация хеша
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("HashPassword | ошибка генерации хеша пароля: %s", err)
	}
	return string(hashedPassword), nil
}

// Проверка, совпадает ли пароль с его хешем
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Сохраняем пользователя в БД
func SaveUserToDB(login string, hashedPassword string) error {
	insertQuery := `INSERT INTO users (login, password) VALUES (?,?)`
	_, err := DB.Exec(insertQuery, login, hashedPassword)
	if err != nil {
		return fmt.Errorf("SaveUserToDB | ошибка сохранения пользователя в базу данных: %s", err)
	}

	log.Printf("Пользователь %s успешно добавлен.", login)
	return nil
}

// Аутентификация пользователя через БД
func AuthenticateUser(login, password string) error {
	// Получение хешированного пароля пользователя из БД
	var hashedPassword string
	query := `SELECT password FROM users WHERE login = ?`
	err := DB.QueryRow(query, login).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrInvalidCredentials
		}
		return err
	}

	// Сравнение введённого пароля с хешем
	err = CheckPassword(hashedPassword, password)
	if err != nil {
		return ErrInvalidCredentials
	}

	log.Printf("Аутентификация пользователя \"%s\" успешна!", login)
	return nil
}

func CheckUserExists(login string) (bool, error) {
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM users WHERE login = ? LIMIT 1);"
	err := DB.QueryRow(checkQuery, login).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("CheckUserExists | ошибка поиска пользователя в бд: %s", err)
	}
	return exists, nil
}

// Функция для регистрации нового пользователя в базе данных
func RegisterUser(login, password string) error {
	// Проверяем, существует ли пользователь с данным логином
	userExists, err := CheckUserExists(login)
	if err != nil {
		return err
	}

	if userExists {
		return ErrUserAlreadyExists
	}

	// Хешируем пароль перед его сохранением в базу данных
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("пользователь \"%s\" ошибка хеширования пароля: %v", login, err)
	}

	// SQL запрос для добавления нового пользователя в базу данных
	query := `INSERT INTO users (login, password) VALUES (?, ?);`
	_, err = DB.Exec(query, login, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("ошибка при добавлении пользователя: %v", err)
	}

	return nil
}
