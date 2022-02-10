package models

import "time"

type Diary struct {
	Id       int
	Title    string
	Body     string
	CreateAt time.Time
}
