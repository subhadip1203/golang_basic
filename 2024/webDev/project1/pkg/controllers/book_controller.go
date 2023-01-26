package controller

import (
	"fmt"
	"net/http"
)

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1 , message: 'get all books'}")
}

func GetBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'get a book'}")
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'Create book' }")
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'Update book' }")
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'Delete book'}")
}
