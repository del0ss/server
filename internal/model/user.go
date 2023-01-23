package model

type User struct {
	ID          int    `json:"ID -"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	EncPassword string `json:"encPassword"`
}
