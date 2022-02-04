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
	Checkinbags   string  `json:"checkinbags"`
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
type Price1 struct{
	Numberofbags string `json:"Number of bags"`
	Price string `json:"price"`
	Dimensions string `json:"dimensions"`
	}
	type Price2 struct{
		Numberofbags string `json:"Number of bags"`
		Price string `json:"price"`
		Dimensions string `json:"dimensions"`
		}
		type Price3 struct{
			Numberofbags string `json:"Number of bags"`
			Price string `json:"price"`
			Dimensions string `json:"dimensions"`
			}





// Init books var as a slice Book struct
var books []Book
var replay []Reply
var a []Price1
var b []Price2
var c []Price3

// Get all bookings
func getBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single booking
func getBooking(w http.ResponseWriter, r *http.Request) {
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

// Add new booking
func createBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Custid = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update booking
func updateBooking(w http.ResponseWriter, r *http.Request) {
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

// Delete booking
func deleteBooking(w http.ResponseWriter, r *http.Request) {
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
func replyback(w http.ResponseWriter,r *http.Request) string{
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _, item := range books {
		if item.Checkinbags == params["checkinbags"] {
			json.NewEncoder(w).Encode(item)
			if(item.Checkinbags=="1"){
				return "10$"
			}
			if(item.Checkinbags=="2"){
				return "20$"
			}
			if(item.Checkinbags=="3"){
				return "30$"
			}
			
		}
	
}
return ""
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
func getprice1(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(a)
	return
		}
		func getprice2(w http.ResponseWriter,r *http.Request){
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(b)
			return
				}
				func getprice3(w http.ResponseWriter,r *http.Request){
					w.Header().Set("Content-Type","application/json")
					json.NewEncoder(w).Encode(c)
					return
						}


// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{Triptype:"One-way",Custid: "John01", Checkinbags:"3", Source: "Florida", Destination: "Alabama", Arrival:"03-04-22"})
	books = append(books, Book{Triptype:"One-way",Custid: "John02", Checkinbags:"2", Source: "New York", Destination: "Tennessee",Arrival:"02-02-22"})
	replay = append(replay, Reply{Custid: "John01", Finalprice:"30$", Luggagecollectiondate: "03-01-22", Courrierpartner: "FED-EX"})
	replay = append(replay, Reply{Custid: "John02", Finalprice:"20$", Luggagecollectiondate: "01-29-22", Courrierpartner: "UPS"})
	a=append(a,Price1{Numberofbags:"One bag only",Price:"10$(Ten Dollars)",Dimensions:"Maximum allowable 22*14 (inches)"})
	b=append(b,Price2{Numberofbags:"Two bags only",Price:"20$(Twenty Dollars)",Dimensions:"Maximum allowable 25*16 (inches)"})
	c=append(c,Price3{Numberofbags:"Three bags only",Price:"30$(Thirty Dollars)",Dimensions: "Maximum allowable 29*18 (inches)"})
	
	


	// Route handles & endpoints
	r.HandleFunc("/bookings", getBookings).Methods("GET")
	r.HandleFunc("/bookings/{custid}", getBooking).Methods("GET")
	r.HandleFunc("/bookings", createBooking).Methods("POST")
	r.HandleFunc("/bookings/{custid}", updateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{custid}", deleteBooking).Methods("DELETE")
	r.HandleFunc("/receipt", getResults).Methods("GET")
	r.HandleFunc("/receipt/{custid}", getResult).Methods("GET")
	// r.HandleFunc("/checkin/{checkinbags}",replyback)
	r.HandleFunc("/luggages/one", getprice1).Methods("GET")
	r.HandleFunc("/luggages/two", getprice2).Methods("GET")
	r.HandleFunc("/luggages/three", getprice3).Methods("GET")



	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

	


