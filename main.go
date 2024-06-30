package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Book struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	AuthorID         int    `json:"author_id"`
	AuthorName       string `json:"author_name"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description,omitempty"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=youruser dbname=yourdb password=yourpassword host=localhost sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page := queryParams.Get("page")
	limit := queryParams.Get("limit")
	sortBy := queryParams.Get("sortBy")
	order := queryParams.Get("order")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	if sortBy == "" {
		sortBy = "id"
	}
	if order == "" {
		order = "asc"
	}

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	sqlQuery := `
		SELECT books.id, books.title, books.author_id, authors.name, books.short_description
		FROM books
		JOIN authors ON books.author_id = authors.id
		ORDER BY books.` + sortBy + ` ` + order + `
		LIMIT $1 OFFSET $2`
	rows, err := db.Query(sqlQuery, limitInt, offset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.AuthorName, &book.ShortDescription)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book Book
	sqlQuery := `
		SELECT books.id, books.title, books.author_id, authors.name, books.short_description, books.full_description
		FROM books
		JOIN authors ON books.author_id = authors.id
		WHERE books.id = $1`
	err = db.QueryRow(sqlQuery, id).Scan(&book.ID, &book.Title, &book.AuthorID, &book.AuthorName, &book.ShortDescription, &book.FullDescription)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
