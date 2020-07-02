package models

import (
	"errors"
	"fmt"
)

// Author represents an author or creator of the book
type Author struct {
	ID        uint   `gorm:"primary_key"`
	Firstname string `json:"FirstName"`
	Lastname  string `json:"LastName"`
}

// TableName change the by-dafault name of the table
func (*Author) TableName() string {
	return "author"
}

// Book entity
type Book struct {
	Isbn     string `gorm:"primary_key" json:"ISBN"`
	Name     string `json:"name"`
	Author   Author `gorm:"foreignkey:AuthorID" json:",omitempty"`
	AuthorID uint   `json:"-"`
}

// TableName change the by-dafault name of the table
func (*Book) TableName() string {
	return "book"
}

// GetBooks return all the books existing in DB
func GetBooks() []Book {
	books := make([]Book, 0)
	GetDB().Preload("Author").Find(&books)
	return books
}

// GetAuthors return all the books existing in DB
func GetAuthors() []Author {
	authors := make([]Author, 0)
	GetDB().Find(&authors)
	return authors
}

// GetAuthor return an Author given its id
func GetAuthor(id int) Author {
	author := Author{}
	GetDB().Find(&author, id)
	return author
}

// GetBook return an Author given its ISBN
func GetBook(id int) Book {
	var book Book
	GetDB().Preload("Author").Find(&book, id)
	return book
}

// InsertAuthor insert a new author in DB
func InsertAuthor(author Author) {
	GetDB().Create(&author)
}

// InsertBook insert a new author in DB
func InsertBook(book Book) error {
	db := GetDB()
	if book.Author == (Author{}) {
		return errors.New("no_author_provided")
	}
	authorTemp := book.Author
	var author Author
	db.Where("firstname = ? AND lastname = ?", authorTemp.Firstname,
		authorTemp.Lastname).First(&author)
	fmt.Println(author)
	if author == (Author{}) {
		return errors.New("author_does_not_exist")
	}
	book.Author = author
	db.Create(&book)
	return nil
}
