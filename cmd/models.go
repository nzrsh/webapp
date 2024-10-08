package main

import "errors"

type FileInfo struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	ModTime string `json:"modTime"`
	IsImage bool   `json:"isImage"`
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var ErrInvalidCredentials = errors.New("неверный логин или пароль")
var ErrUserAlreadyExists = errors.New("пользователь уже существует")
var ErrEmptyLogin = errors.New("логин не может быть пустым")
var ErrEmptyPassword = errors.New("пароль не может быть пустым")

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserData struct {
	Login string `json:"login"`
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
