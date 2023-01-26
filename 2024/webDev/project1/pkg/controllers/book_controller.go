package controller

import (
	"fmt"
	"net/http"
)

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func GetBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}
