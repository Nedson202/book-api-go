package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"database/sql"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setupDatabaseConnection() {
	dbURL, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", dbURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)
}

func main() {
	setupDatabaseConnection()

	router := mux.NewRouter()
	port := ":7000"

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
			logFatal(err)
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

func getBooks(w http.ResponseWriter, req *http.Request) {
	var bookCopy book

	books = []book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bookCopy.ID, &bookCopy.Title, &bookCopy.Author, &bookCopy.Year)
		logFatal(err)

		books = append(books, bookCopy)
	}

	json.NewEncoder(w).Encode(books)
	log.Println("Get all books")
}

func getBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	parsedID, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == parsedID {
			json.NewEncoder(w).Encode(&book)
		}
	}
	log.Println("Get a book", params)
}

func addBook(w http.ResponseWriter, req *http.Request) {
	var newBook book

	json.NewDecoder(req.Body).Decode(&newBook)
	newBook.ID = len(books) + 1

	books = append(books, newBook)

	json.NewEncoder(w).Encode(newBook)
	log.Println("Add a new book", books, newBook)
}

func updateBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	parsedID, _ := strconv.Atoi(params["id"])

	var updateValues book

	updateValues.ID = parsedID

	json.NewDecoder(req.Body).Decode(&updateValues)

	log.Println("its length", len(books)-1)

	for index, book := range books {
		if book.ID == parsedID {
			books[index] = updateValues
		}
	}

	json.NewEncoder(w).Encode(updateValues)
	log.Println("Update a book", updateValues, parsedID)
}

func removeBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	parseID, _ := strconv.Atoi(params["id"])

	for index, book := range books {
		if book.ID == parseID {
			books = append(books[:index], books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
	log.Println("Remove a book", books, parseID)
}
