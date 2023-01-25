package controller

import (
	"fmt"
	"net/http"
)

func getAllBooks(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func getBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func createBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func updateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1}")
}
