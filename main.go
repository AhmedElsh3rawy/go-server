package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"main/handler"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	// handle user routes
	router.HandleFunc("/users", handler.AddUser).Methods("POST")
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetUserById)

	// Start the server on port 8080
	fmt.Println("Server starting on localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
