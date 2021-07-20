package main 

import {
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
}

//Book struct 

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string`json:"lastname"`
}

//slice to hold the books 
books := []Book

//Get all books 
func getBooks(w http.ResponseWriter, r *http,Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a single book 
func getBook(w http.ResponseWriter, r *http,Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params 
	//Loop through the books and find the one with the ID from the params 
	for _, item :- range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&book{})
}

//Add new book 
func addNewBook(w http.ResponseWriter, r *http,Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode($book)
	book.ID = strvonv.Itoa(rand.Intn(100000000)) //Mock ID - Not safe 
	books = append(books, book )
	json.NewEncoder(w).Encode(book)
}

//Update book 
func updateBook(w http.ResponseWriter, r *http,Request) {
	w.Header().Set("Content-Type", "application/json")
	params :- mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index + 1:]...)
			var book Book_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}