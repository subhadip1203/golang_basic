package route

import (
	"github.com/gorilla/mux"
	controller "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/controllers"
)

var BookStaoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controller.GetAllBooks).Methods(("GET"))
	router.HandleFunc("/book/{bookId}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}
