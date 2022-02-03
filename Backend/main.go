package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	Triptype     string  `json:"triptype"`
	checkinbags   string  `json:"checkinbags"`
	Source string  `json:"source"`
	Destination string `json:"destination"`
	Arrival string `json:"arrival"`
	Custid string `json:"custid"`
}
type Reply struct{
Custid string `json:"custid"`
Finalprice string `json:"finalprice"`
Luggagecollectiondate string `json:"luggagecollectiondate"`
Courrierpartner string `json:"courrierpartner"`
}

// Author struct


// Init books var as a slice Book struct
var books []Book
var replay []Reply

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.Custid == params["custid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Custid = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Custid == params["custid"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Custid = params["custid"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Custid == params["custid"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func replyback(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _, item := range books {
		if item.Custid == params["custid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	
}
}
func getResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(replay)
}
func getResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range replay {
		if item.Custid == params["custid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Reply{})
}


// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{Triptype:"One-way",Custid: "John01", checkinbags:"3", Source: "Florida", Destination: "Alabama", Arrival:"03-04-22"})
	books = append(books, Book{Triptype:"One-way",Custid: "John02", checkinbags:"2", Source: "New York", Destination: "Tennessee",Arrival:"02-02-22"})
	replay = append(replay, Reply{Custid: "John01", Finalprice:"35$", Luggagecollectiondate: "03-01-22", Courrierpartner: "FED-EX"})
	replay = append(replay, Reply{Custid: "John02", Finalprice:"30$", Luggagecollectiondate: "01-29-22", Courrierpartner: "UPS"})

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{custid}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{custid}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{custid}", deleteBook).Methods("DELETE")
	r.HandleFunc("/reply", getResults).Methods("GET")
	r.HandleFunc("/reply/{custid}", getResult).Methods("GET")


	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

	


