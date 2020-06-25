package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/macperez/binovoo/models"
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
