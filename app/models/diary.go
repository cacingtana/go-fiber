package models

import "time"

type Diary struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	CreateAt time.Time `json:"createAt"`
}
