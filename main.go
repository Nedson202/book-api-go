package main

import (
	// "fmt"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []book

func main() {
	router := mux.NewRouter()
	port := ":7000"

	books = append(books,
		book{ID: 1, Title: "Golang pointers", Author: "Mr Golang", Year: "2000"},
		book{ID: 2, Title: "Golang pointers", Author: "Mr Golang", Year: "2000"},
		book{ID: 3, Title: "Golang pointers", Author: "Mr Golang", Year: "2000"},
		book{ID: 4, Title: "Golang pointers", Author: "Mr Golang", Year: "2000"},
		book{ID: 5, Title: "Golang pointers", Author: "Mr Golang", Year: "2000"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Start Server
	func() {
		log.Println("Starting Server on port", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	handleShutdown(server)
}

// Handle graceful shutdown
func handleShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
	log.Println("Get all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parsedID, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == parsedID {
			json.NewEncoder(w).Encode(&book)
		}
	}
	log.Println("Get a book", params)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook book

	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = len(books) + 1

	books = append(books, newBook)

	json.NewEncoder(w).Encode(newBook)
	log.Println("Add a new book", books, newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parsedID, _ := strconv.Atoi(params["id"])

	var updateValues book

	updateValues.ID = parsedID

	json.NewDecoder(r.Body).Decode(&updateValues)

	log.Println("its length", len(books)-1)

	for index, book := range books {
		if book.ID == parsedID {
			books[index] = updateValues
		}
	}

	json.NewEncoder(w).Encode(updateValues)
	log.Println("Update a book", updateValues, parsedID)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parseID, _ := strconv.Atoi(params["id"])

	for index, book := range books {
		if book.ID == parseID {
			books = append(books[:index], books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
	log.Println("Remove a book", books, parseID)
}
