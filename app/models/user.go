package models

import "time"

type UserCredential struct {
	Email  string
	Status string
}

type Role struct {
	Id   uint
	Name string
}

type User struct {
	Id       uint
	Email    string
	Password string
	Level    []Role
	CreateAt time.Time
}
