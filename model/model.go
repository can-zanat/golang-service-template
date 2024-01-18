package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
