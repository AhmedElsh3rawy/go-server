package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"main/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
	router := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.HandleFunc("/", homeHandler)
	// handle user routes
	router.HandleFunc("/users", handlers.GetUsers)
	router.HandleFunc("/users/{id}", handlers.FindUserById)

	// Start the server on port 8080
	PORT := ":" + os.Getenv("PORT")
	fmt.Printf("Server starting on localhost%s\n", PORT)
	if err := http.ListenAndServe(PORT, router); err != nil {
		panic(err)
	}
}
