package models

import (
	"github.com/Adityap-3055/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

// db variable here gets the data from .GetDB()
//
//	and auto migrate to BOOK struct
var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// helps to initalize our database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// in order to make changes to the database it has
// to go through this models using these model functions

func (b *Book) CreateBook() *Book {
	// query will be writen by gorm, we only need to write the functions
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
