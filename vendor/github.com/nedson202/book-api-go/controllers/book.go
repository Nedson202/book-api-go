package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/nedson202/book-api-go/config"
	"github.com/nedson202/book-api-go/models"
	"github.com/nedson202/book-api-go/repository/book"
)

var books []models.BookSchema

// Controller struct
type Controller struct{}

var bookRepo = bookRepository.BookRepository{}

// GetBooks controller
func (c Controller) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var book models.BookSchema

		books = bookRepo.GetBooks(book)

		config.RespondWithJSON(w, http.StatusOK, books)
	}
}

// GetBook controller to retrieve a book from db
func (c Controller) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		id, err := strconv.Atoi(params["id"])
		config.LogFatal(err)

		var book models.BookSchema

		book = bookRepo.GetBook(book, id)

		if book.ID == 0 {
			config.RespondWithError(w, http.StatusNotFound, "No match found for book requested")
		} else {
			config.RespondWithJSON(w, http.StatusOK, book)
		}
	}
}

// AddBook controller to add a book to db
func (c Controller) AddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		var newBook models.BookSchema

		json.NewDecoder(req.Body).Decode(&newBook)

		newBook = bookRepo.AddBook(newBook)
		config.RespondWithJSON(w, http.StatusCreated, newBook)
	}
}

// UpdateBook controller to update book record on db
func (c Controller) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		params := mux.Vars(req)
		parsedID, _ := strconv.Atoi(params["id"])

		var updateValues models.BookSchema

		updateValues.ID = parsedID

		json.NewDecoder(req.Body).Decode(&updateValues)

		rowsUpdated := bookRepo.UpdateBook(updateValues, parsedID)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

// RemoveBook controller to delete book from db
func (c Controller) RemoveBook() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		params := mux.Vars(req)

		id, err := strconv.Atoi(params["id"])
		config.LogFatal(err)

		rowsDeleted := bookRepo.RemoveBook(id)
		if rowsDeleted == 1 {
			config.RespondWithJSON(w, http.StatusOK,
				config.Payload{Error: true, Message: "Book successfully deleted"})
		} else {
			config.RespondWithError(w, http.StatusInternalServerError,
				"An error occurred trying to delete book")
		}
	}
}
