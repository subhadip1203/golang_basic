package route

import (
	"github.com/gorilla/mux"
)

var bookStaoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.getAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBook).Methods("GET")
	router.HandleFunc("/book", controllers.createBook).Methods("POST")
	router.HandleFunc("/book", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")
}
