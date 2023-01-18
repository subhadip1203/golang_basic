package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    ` json: "id" `
	Isbn     string    ` json: "isbn" `
	Title    string    `json: "title" `
	Director *Director `json: "director" `
}

type Director struct {
	FirstName string ` json: "firstname" `
	LastName  string ` json: "lastname" `
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ping).Methods("GET")
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}

// function for route : "/" GET
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{status:1}")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {}
func getMovie(w http.ResponseWriter, r *http.Request)     {}
func createMovies(w http.ResponseWriter, r *http.Request) {}
func updateMovies(w http.ResponseWriter, r *http.Request) {}
func deleteMovies(w http.ResponseWriter, r *http.Request) {}
