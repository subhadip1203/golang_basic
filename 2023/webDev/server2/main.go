package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/form", aboutForm)

	fmt.Println("starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "about us page")
}

func aboutForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseForm error : %v", err)
		return
	}
	fmt.Fprintf(w, "form parse successfully \n")
	name := r.FormValue("name")
	address := r.FormValue(("address"))
	fmt.Fprintf(w, "name= %s \n", name)
	fmt.Fprintf(w, "address= %s \n", address)
}
