package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/macperez/books/controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	router.HandleFunc("/author/{id:[0-9]+}", controllers.GetAuthor).Methods("GET")
	router.HandleFunc("/author", controllers.CreateNewAuthor).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id:[0-9]+}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", controllers.CreateNewBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
