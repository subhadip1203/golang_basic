package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string ` json:"name" `
	Author      string ` json:"author" `
	Publication string ` json:"publication" `
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		return nil
	}
	return b
}

func (b *Book) UpdateBook() *Book {
	db.First(&b)
	db.Save(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) *Book {
	var GetBook Book
	db.First(&GetBook, id)
	return &GetBook
}

func DeleteBook(id int64) *Book {
	var DeletedBook Book
	db.Delete(&DeletedBook, 10)
	return &DeletedBook
}
