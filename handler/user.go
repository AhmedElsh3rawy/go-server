package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AhmedElsh3rawy/go-server/database"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

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

func AddUser(w http.ResponseWriter, r *http.Request) {}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
