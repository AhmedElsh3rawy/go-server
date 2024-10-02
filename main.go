package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AhmedElsh3rawy/go-server/handler"
	"github.com/AhmedElsh3rawy/go-server/middleware"
	"github.com/joho/godotenv"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, stranger..."))
}

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /", getHello)

	stack := middleware.CreateStack(
		middleware.Logging)

	router.HandleFunc("GET /users", handler.GetUsers)
	router.HandleFunc("GET /users/{id}", handler.GetUser)
	router.HandleFunc("POST /users", handler.AddUser)
	router.HandleFunc("PATCH /users/{id}", handler.UpdateUser)
	router.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("[server]: starting on localhost:8080")
	fmt.Println(`
    ░██████╗░░█████╗░  ░█████╗░██████╗░██╗
    ██╔════╝░██╔══██╗  ██╔══██╗██╔══██╗██║
    ██║░░██╗░██║░░██║  ███████║██████╔╝██║
    ██║░░╚██╗██║░░██║  ██╔══██║██╔═══╝░██║
    ╚██████╔╝╚█████╔╝  ██║░░██║██║░░░░░██║
    ░╚═════╝░░╚════╝░  ╚═╝░░╚═╝╚═╝░░░░░╚═╝`)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
