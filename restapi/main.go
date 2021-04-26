package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//book struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Author
type Author struct {
	Firstname string `json:"firstname`
	Lastname string `json:"lastname`
}

//init books var a slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // Get any params
	//Loop through books and find with id 
	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book 
	_=json.NewDecoder(r.Body).Decode(&book)
	book.ID= strconv.Itoa(rand.Intn(10000000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books =append(books[:index],books[index+1:]...)
			var book Book 
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item  := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("asd")
	//Init router
	r := mux.NewRouter()

	//mockdata
	books =append(books, Book{ID:"1",Isbn:"234233",Title:"Book One", Author: &Author{Firstname:"john",Lastname:"doe"}})
	books =append(books, Book{ID:"2",Isbn:"54354233",Title:"Book Two", Author: &Author{Firstname:"john",Lastname:"doe"}})
	books =append(books, Book{ID:"3",Isbn:"342233",Title:"Book Three", Author: &Author{Firstname:"john",Lastname:"doe"}})


	//Route handlers/ endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",r ))
}