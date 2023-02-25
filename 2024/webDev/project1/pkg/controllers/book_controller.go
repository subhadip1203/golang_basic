package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/models"
	"github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/utils"
)

func GetAllBooks(res http.ResponseWriter, req *http.Request) {

	//DB query to find data & handling DB search error
	books, err := model.GetAllBooks()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	// Handling JSON conversionn and error
	jsonBooks, err := json.Marshal(books)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("{status:0}"))
		return
	}

	res.Header().Set("content-type", "appilaction/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBooks)

}

func GetBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]

	// paring param bookid to int
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("{status:0}"))
		return
	}

	// searching book in DB & handling error
	book, err := model.GetBookByID(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	// Handling JSON conversionn and error
	jsonBook, err := json.Marshal(book)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("{status:0}"))
		return
	}

	// if all error free , sending respoands
	res.Header().Set("content-type", "appilaction/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBook)

}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	// Saving book in DB
	book := &model.Book{}
	utils.ParseBody(req, book)
	newBook, err := book.CreateBook()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	// Handling JSON conversionn and error
	jsonBook, err := json.Marshal(newBook)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("{status:0}"))
		return
	}

	// if all error free , sending respoands
	res.Header().Set("content-type", "appilaction/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBook)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'Update book' }")
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{status:1, message: 'Delete book'}")
}
