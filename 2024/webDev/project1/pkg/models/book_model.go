package model

import (
	"errors"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string ` json:"name" `
	Author      string ` json:"author" `
	Publication string ` json:"publication" `
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(&b)
	if result.Error != nil {
		return nil, errors.New("Book not saved")
	}
	return b, nil
}

func (b *Book) UpdateBook() *Book {
	db.First(&b)
	db.Save(&b)
	return b
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	result := db.Find(&Books)
	if result.Error != nil {
		return nil, errors.New("no data found")
	}
	return Books, nil
}

func GetBookByID(id int64) (*Book, error) {
	var GetBook Book
	result := db.First(&GetBook, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return &GetBook, errors.New("no data found")
	}
	return &GetBook, nil
}

func DeleteBook(id int64) (*Book, error) {
	var DeletedBook Book
	result := db.Delete(&DeletedBook, id)
	if result.Error != nil {
		return nil, errors.New("Book not Deleted")
	}
	return &DeletedBook, nil
}
