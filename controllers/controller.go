package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/macperez/books/models"
)

func message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetBooks calls the model to bring all the books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	data := models.GetBooks()
	resp := message(true, "success")
	resp["books"] = data
	respond(w, resp)
}

// GetAuthors calls the model to bring all the authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	data := models.GetAuthors()
	resp := message(true, "success")
	resp["authors"] = data
	respond(w, resp)
}

// GetAuthor calls the model to bring one specific author
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.GetAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Author not in database\n"))

	} else {
		resp := message(true, "success")
		resp["author"] = data
		respond(w, resp)
	}

}

// GetBook calls the model to bring one specific book
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.GetBook(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Book not in database\n"))

	} else {
		resp := message(true, "success")
		resp["book"] = data
		respond(w, resp)
	}
}

// CreateNewAuthor calls the model to create one author in database
func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var author models.Author
	json.Unmarshal(reqBody, &author)
	models.InsertAuthor(author)
	GetAuthors(w, r)

}

// CreateNewBook calls the model to create one book in database
func CreateNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var book models.Book
	json.Unmarshal(reqBody, &book)
	err := models.InsertBook(book)
	if err != nil {
		switch err.Error() {
		case "author_does_not_exist":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Author not in database\n"))
		case "no_author_provided":
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - No Author provided\n"))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal server error"))
		}
	}

	GetBooks(w, r)

}
