package models

import (
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
	Author   Author `gorm:"foreignkey:AuthorID" json:"author,omitempty"`
	AuthorID uint
}

// TableName change the by-dafault name of the table
func (*Book) TableName() string {
	return "book"
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
	authorsQuery := `SELECT id, firstname, lastname FROM binovoo.author`
	authors := make([]Author, 0)
	GetDB().Raw(authorsQuery).Scan(&authors)
	return authors
}

// InsertAuthor insert a new author in DB
func InsertAuthor(author Author) {
	GetDB().Create(author)
}

// InsertBook insert a new author in DB
func InsertBook(book Book) {
	db := GetDB()
	fmt.Printf("\n%+v\n", book)
	if book.Author == (Author{}) {
		fmt.Println("NULO")
	} else {
		authorTemp := book.Author
		var author Author
		db.Where("firstname = ? AND lastname = ?", authorTemp.Firstname,
			authorTemp.Lastname).First(&author)
		book.Author = author
		fmt.Println(author)
		db.Create(book)

	}

	// GetDB().Create(book)
}

/*
func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}

	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}
*/
