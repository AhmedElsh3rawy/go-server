package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AhmedElsh3rawy/go-server/database"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db := database.Db

	rows, err := db.Query("SELECT username, email FROM users;")
	if err != nil {
		log.Printf("error querying users table %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Username, &u.Email)
		if err != nil {
			log.Printf("error while scannig rows %v", err)
			return
		}

		users = append(users, u)
	}

	j, err := json.Marshal(users)
	if err != nil {
		log.Printf("error marshalling users into json %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

func GetUser(w http.ResponseWriter, r *http.Request) {}
// get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Printf("Error while converting string to number: %v", err)
		return
	}

	db := database.Db
	var user User
	e := db.QueryRow("SELECT username, email FROM users WHERE id = $1;", id).
		Scan(&user.Username, &user.Email)
	if e != nil {
		fmt.Printf("error while scannig row: %v", e)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		log.Printf("error marshalling users into json %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
