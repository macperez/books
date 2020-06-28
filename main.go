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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	myRouter.HandleFunc("/authors", controllers.GetAuthors).Methods("GET")
	myRouter.HandleFunc("/author", controllers.CreateNewAuthor).Methods("POST")
	myRouter.HandleFunc("/book", controllers.CreateNewBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
