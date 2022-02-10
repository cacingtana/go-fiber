package models

import "time"

type UserCredential struct {
	Email  string `json:"id"`
	Status string `json:"status"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id       int       `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Level    int       `json:"level"`
	CreateAt time.Time `json:"createAt"`
}
