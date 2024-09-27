package main

import "errors"

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var ErrInvalidCredentials = errors.New("неверный логин или пароль")
var ErrUserAlreadyExists = errors.New("пользователь уже существует")

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Product struct {
	ID    int     `json:"id"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

/* JSON MODEL
{
	"type":"cpu",
	"name":"name",
	"price":price
}


*/

var CreateTableUsersQuery string = `
CREATE TABLE IF NOT EXISTS "users" (
	"id"	INTEGER NOT NULL UNIQUE,
	"login"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

var CreateTableProductsQuery string = `
CREATE TABLE IF NOT EXISTS "products" (
	"id"	INTEGER NOT NULL UNIQUE,
	"type"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"price"	REAL NOT NULL,
	PRIMARY KEY("id")
);
`
