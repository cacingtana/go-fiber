package models

import "time"

type UserCredential struct {
	Email  string
	Status string
}

type Role struct {
	Id   int
	Name string
}

type User struct {
	Id       int
	Email    string
	Password string
	Level    int
	CreateAt time.Time
}
