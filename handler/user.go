package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AhmedElsh3rawy/go-server/database"
	"github.com/AhmedElsh3rawy/go-server/users"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreatedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := database.ConnectDatabase()
	queries := users.New(db)

	users, err := queries.GetUsers(ctx)
	if err != nil {
		http.Error(w, "Error quering database", http.StatusInternalServerError)
		return
	}

	var us []User
	for _, e := range users {
		var u User
		u.ID = e.ID
		u.Username = e.Username.String
		u.Email = e.Email.String
		us = append(us, u)
	}

	jsonData, err := json.Marshal(us)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Printf("Error while converting string to number: %v", err)
		return
	}

	ctx := context.Background()
	db := database.ConnectDatabase()
	queries := users.New(db)
	user, err := queries.GetUser(ctx, int32(id))

	if err != nil {
		http.Error(w, "Error quering database", http.StatusInternalServerError)
		return
	}

	var u User
	u.ID = user.ID
	u.Username = user.Username.String
	u.Email = user.Email.String

	jsonData, err := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

// add user
func AddUser(w http.ResponseWriter, r *http.Request) {
	var u CreatedUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println("error while decoding request body")
		return
	}

	w.Write([]byte("User added"))

}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
