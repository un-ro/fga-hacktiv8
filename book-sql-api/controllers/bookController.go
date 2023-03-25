package controllers

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver for postgres
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/fga?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Create
func CreateBook() {
	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/book_api?sslmode=disable")
}
