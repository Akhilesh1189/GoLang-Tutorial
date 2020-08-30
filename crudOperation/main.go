package main

import (
    "encoding/json"
    "log"
    "net/http"
    "math/rand"
   "strconv"
    "github.com/gorilla/mux"
  )
  
 //Book Struct (Model)
 type Book struct {
     ID string `json:"id"`
     Isbn string `json:"isbn"`
     Title string `json:"title"`
     Author *Author `json:"author"`

 }
 // Author Struct
 
 type Author struct {
     FirstName string `json:"firstname"`
     LastName string `json:"lastname"`
 }
 // Init books var as a slice Book struct
 var books []Book
 // Get All Books

 func getBooks(w http.ResponseWriter, r *http.Request) {
     w.Header().Set("Content-Type","application/json")
     json.NewEncoder(w).Encode(books)

 }
 // Get Single Book
 func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for _, item := range books {
        if item.ID == params["id"] {
        json.NewEncoder(w).Encode(item)
            return
        }

    }
 json.NewEncoder(w).Encode(&Book{})
}
// Create New Book
func creatBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID 
    books = append(books, book)
    json.NewEncoder(w).Encode(book)

}
//
func updateBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for index,item := range books {
        if item.ID == params["id"] {
        books =append(books[:index],books[index+1:]...)
        // var bool Book
        //  _ = json.NewDecoder(r.Body).Decode(&book)
        //  book.ID = strconv.Itoa(rand.Intn(10000000))
        //  books = append(books, book)
        //  json.NewDecoder(w).Encode(book)
        // return
        }
    }
   // json.NewDecoder(w).Encode(books)
}
func deleteBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for index,item := range books {
        if item.ID == params["id"] {
        books =append(books[:index],books[index+1:]...)
        break
        }
    }
    json.NewEncoder(w).Encode(books)
}
func main() {
        r := mux.NewRouter()

    books = append(books, Book{ID:"1",Isbn:"448743",Title:"Book One",Author:&Author{FirstName:"Akhilesh",LastName:"Pandey"}})
    books = append(books, Book{ID:"2",Isbn:"448744",Title:"Book Two",Author:&Author{FirstName:"Pankaj",LastName:"vasitha"}})
   
   // json.NewEncoder(w).Encode(Book)
    //Route Handlers / Endpoints
    r.HandleFunc("/api/books",getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
    r.HandleFunc("/api/books",creatBooks).Methods("POST")
    r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
    r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000",r))
}