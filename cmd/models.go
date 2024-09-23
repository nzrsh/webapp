package main

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Product struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var CreateTableUsersQuery string = `
CREATE TABLE "users" (
	"id"	INTEGER NOT NULL UNIQUE,
	"login"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

var CreateTableProductsQuery string = `
CREATE TABLE "products" (
	"id"	INTEGER NOT NULL UNIQUE,
	"type"	TEXT NOT NULL,
	"name"	TEXT NOT NULL,
	"price"	REAL NOT NULL,
	PRIMARY KEY("id")
);
`
