package main

import (
    "encoding/json"
    "log"
    "net/http"
    //"math/rand"
   //"strconv"
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

}
//
func updateBooks(w http.ResponseWriter, r *http.Request) {

}
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}
func main() {
    // fmt.Println("Hello World!")
    // for i := 0; i < 10; i++ {
    //     println(rand.Intn(25))
    //   }
    r := mux.NewRouter()

    books = append(books, Book{ID:"1",Isbn:"448743",Title:"Book One",Author:&Author{FirstName:"Akhilesh",LastName:"Pandey"}})
    books = append(books, Book{ID:"2",Isbn:"448744",Title:"Book Two",Author:&Author{FirstName:"Pankaj",LastName:"vasitha"}})
   
    //Route Handlers / Endpoints
    r.HandleFunc("/api/books",getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
    r.HandleFunc("/api/books",creatBooks).Methods("POST")
    r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
    r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000",r))
}