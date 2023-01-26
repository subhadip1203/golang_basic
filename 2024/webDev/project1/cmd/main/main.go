package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	route "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	route.BookStaoreRoutes(r)
	r.HandleFunc("/", ping)
	fmt.Println("starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

// function for route : "/" GET
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{status:1}")
}
