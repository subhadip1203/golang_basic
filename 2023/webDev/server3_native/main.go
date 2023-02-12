package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeContoller(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(`{"message": "POST method requested"}`))
	default:
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func aboutContoller(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"message": "About Route API"}`))

}

func main() {
	http.HandleFunc("/", homeContoller)
	http.HandleFunc("/about", aboutContoller)
	fmt.Println("starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
