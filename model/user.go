package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"firstName"`
	Email string `json:"email"`
}
