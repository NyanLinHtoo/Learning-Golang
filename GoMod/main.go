package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	log.Println("Server is starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
