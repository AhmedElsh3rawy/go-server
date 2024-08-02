package main

import (
	"fmt"
	"net/http"

	"github.com/AhmedElsh3rawy/go-server/database"
	"github.com/AhmedElsh3rawy/go-server/handler"
	"github.com/AhmedElsh3rawy/go-server/middleware"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, stranger..."))
}

func main() {
	database.ConnectDatabase()

	router := http.NewServeMux()

	router.HandleFunc("GET /", getHello)

	stack := middleware.CreateStack(
		middleware.Logging)

	router.HandleFunc("GET /users", handler.GetUsers)

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
	// database.CloseDatabase()
}
