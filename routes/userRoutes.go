package routes

import (
	"github.com/gorilla/mux"

	"main/handlers"
)

func UserRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetUsers)
	return router
}
