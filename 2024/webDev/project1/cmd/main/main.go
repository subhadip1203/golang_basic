package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/config"
	route "github.com/subhadip1203/golang_basic/2024/webDev/project1/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ping)
	route.BookStaoreRoutes(r)
	config.ConnectDB()
	fmt.Println("starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}

// function for route : "/" GET
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{status:1}")
}
