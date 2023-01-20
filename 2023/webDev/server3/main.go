package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    ` json: "id" `
	ISBN     string    ` json: "isbn" `
	Title    string    `json: "title" `
	Director *Director `json: "director" `
}

type Director struct {
	FirstName string ` json: "firstname" `
	LastName  string ` json: "lastname" `
}

var movies []Movie

func main() {

	// creating Fake movies
	movies = append(movies, Movie{ID: "1", ISBN: "abcd1", Title: "Title 1", Director: &Director{FirstName: "Director 1 first", LastName: "Director 1 last"}})
	movies = append(movies, Movie{ID: "2", ISBN: "abcd2", Title: "Title 2", Director: &Director{FirstName: "Director 2 first", LastName: "Director 2 last"}})

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

// get all the movies details
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get a perticular movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {}
func updateMovies(w http.ResponseWriter, r *http.Request) {}

// Delete a movie
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
