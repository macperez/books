package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

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
		fmt.Print(err)
	}
	//GetBooks(w, r)

}
