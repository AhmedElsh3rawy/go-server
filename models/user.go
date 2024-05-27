package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	DoB       string `json:"dob"`
}

var Users = []User{
	{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		DoB:       "2000-01-08",
	}, {
		ID:        2,
		FirstName: "Mary",
		LastName:  "Jane",
		Email:     "maryjane@example.com",
		DoB:       "2002-04-06",
	},
}
