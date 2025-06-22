package models

type User struct {
	Username string
	Password string
	Role     string // "doctor" or "receptionist"
}

var Users = []User{
	{
		Username: "doctor1",
		Password: "pass123",
		Role:     "doctor",
	},
	{
		Username: "reception1",
		Password: "pass123",
		Role:     "receptionist",
	},
}
