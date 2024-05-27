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

func GetUserById(w http.ResponseWriter, r *http.Request) {
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
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
