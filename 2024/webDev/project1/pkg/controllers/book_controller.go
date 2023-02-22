package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/models"
)

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	books := model.GetAllBooks()
	jsonBooks, err := json.Marshal(books)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("{status:0}"))
		return
	} else {
		res.Header().Set("content-type", "appilaction/json")
		res.WriteHeader(http.StatusOK)
		res.Write(jsonBooks)
	}
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
