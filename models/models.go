package models

// Author represents an author or creator of the book
type Author struct {
	Firstname string `json:"FirstName"`
	Lastname  string `json:"LastName"`
}

// Book entity
type Book struct {
	Isbn   string   `json:"ISBN"`
	Name   string   `json:"name"`
	Author []Author `json:"author,omitempty"`
}

// GetBooks return all the books existing in DB
func GetBooks() []Book {
	booksQuery := `SELECT isbn, name FROM binovoo.book`
	books := make([]Book, 0)
	GetDB().Raw(booksQuery).Scan(&books)
	return books
}

// GetAuthors return all the books existing in DB
func GetAuthors() []Author {
	authorsQuery := `SELECT firstname, lastname FROM binovoo.author`
	authors := make([]Author, 0)
	GetDB().Raw(authorsQuery).Scan(&authors)
	return authors
}
