package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"main/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal("Error while converting")
		}

		for _, user := range models.Users {
			if user.ID == id {

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		fmt.Fprintf(w, "Not found")
	} else if r.Method == http.MethodPut {
		// handle put
	} else if r.Method == http.MethodDelete {
		// handle delete
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
}
