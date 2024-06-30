package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Book struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	AuthorIDs        []int    `json:"author_ids"`
	AuthorNames      []string `json:"author_names"`
	ShortDescription *string  `json:"short_description"`
	FullDescription  *string  `json:"full_description,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=USER dbname=fector_go password=PWD host=localhost sslmode=disable"
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
	authorIDs := queryParams.Get("authorIDs")

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
		SELECT books.id, books.title, books.short_description, books.full_description, array_agg(authors.id), array_agg(authors.name)
		FROM books
		JOIN book_authors ON books.id = book_authors.book_id
		JOIN authors ON book_authors.author_id = authors.id`

	if authorIDs != "" {
		authorIDList := strings.Split(authorIDs, ",")
		sqlQuery += " WHERE authors.id IN (" + strings.Join(authorIDList, ",") + ")"
	}

	sqlQuery += `
		GROUP BY books.id
		ORDER BY books.` + sortBy + ` ` + order + `
		LIMIT $1 OFFSET $2`

	rows, err := db.Query(sqlQuery, limitInt, offset)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		var shortDescription sql.NullString
		var fullDescription sql.NullString
		var authorIDs string
		var authorNames string
		err := rows.Scan(&book.ID, &book.Title, &shortDescription, &fullDescription, &authorIDs, &authorNames)
		if err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
			return
		}
		if shortDescription.Valid {
			book.ShortDescription = &shortDescription.String
		}
		if fullDescription.Valid {
			book.FullDescription = &fullDescription.String
		}
		if err := parseArray(authorIDs, &book.AuthorIDs); err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
			return
		}
		if err := parseArray(authorNames, &book.AuthorNames); err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Bad Request", Error: err.Error()})
		return
	}

	var book Book
	var shortDescription sql.NullString
	var fullDescription sql.NullString
	var authorIDs string
	var authorNames string
	sqlQuery := `
		SELECT books.id, books.title, books.short_description, books.full_description, array_agg(authors.id), array_agg(authors.name)
		FROM books
		JOIN book_authors ON books.id = book_authors.book_id
		JOIN authors ON book_authors.author_id = authors.id
		WHERE books.id = $1
		GROUP BY books.id`
	err = db.QueryRow(sqlQuery, id).Scan(&book.ID, &book.Title, &shortDescription, &fullDescription, &authorIDs, &authorNames)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Not Found"})
		} else {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
		}
		return
	}
	if shortDescription.Valid {
		book.ShortDescription = &shortDescription.String
	}
	if fullDescription.Valid {
		book.FullDescription = &fullDescription.String
	}
	if err := parseArray(authorIDs, &book.AuthorIDs); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
		return
	}
	if err := parseArray(authorNames, &book.AuthorNames); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal Server Error", Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func parseArray(arrayString string, target interface{}) error {
	arrayString = strings.Trim(arrayString, "{}")
	elements := strings.Split(arrayString, ",")
	switch v := target.(type) {
	case *[]int:
		*v = make([]int, len(elements))
		for i, elem := range elements {
			val, err := strconv.Atoi(strings.TrimSpace(elem))
			if err != nil {
				return err
			}
			(*v)[i] = val
		}
	case *[]string:
		*v = make([]string, len(elements))
		for i, elem := range elements {
			(*v)[i] = strings.Trim(elem, `"`)
		}
	default:
		return fmt.Errorf("unsupported type: %T", target)
	}
	return nil
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
