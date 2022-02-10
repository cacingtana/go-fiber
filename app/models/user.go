package models

import "time"

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
