package models

import "time"

type Diary struct {
	Id       uint
	Title    string
	Body     string
	CreateAt time.Time
}
