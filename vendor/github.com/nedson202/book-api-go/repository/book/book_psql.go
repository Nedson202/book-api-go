package bookRepository

import (
	"log"
	"database/sql"

	"github.com/nedson202/book-api-go/config"
	"github.com/nedson202/book-api-go/models"
	"github.com/nedson202/book-api-go/driver"
)

type BookRepository struct{}

var books []models.BookSchema

var db *sql.DB 

func (b BookRepository) GetBooks(book models.BookSchema) []models.BookSchema {
	db = driver.DB
	rows, err := db.Query("select * from books")
	config.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		config.LogFatal(err)

		books = append(books, book)
	}

	return books
}

func (b BookRepository) GetBook(book models.BookSchema, id int) (models.BookSchema) {
	db = driver.DB
	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	switch {
		case err == sql.ErrNoRows:
			log.Printf("Book not found")
		case err != nil:
			config.LogFatal(err)
		default:
			return book
	}

	return book
}

func (b BookRepository) AddBook(newBook models.BookSchema) models.BookSchema {
	db = driver.DB
	err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id;",
	newBook.Title, newBook.Author, newBook.Year).Scan(&newBook.ID)

	config.LogFatal(err)
	return newBook
}

func (b BookRepository) UpdateBook(updateValues models.BookSchema, id int) int64 {
	db = driver.DB
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 Returning id;",
		updateValues.Title, updateValues.Author, updateValues.Year, id)

	rowsUpdated, err := result.RowsAffected()
	config.LogFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(id int) int64 {
	db = driver.DB
	result, err := db.Exec("delete from books where id = $1", id)

	rowsDeleted, err := result.RowsAffected()
	config.LogFatal(err)

	return rowsDeleted
}
