package response

import "time"

type diaryResponse struct {
	Id       int       `json:"Id"`
	Title    string    `json:"Title"`
	Body     string    `json:"Body"`
	CreateAt time.Time `json:"createAt"`
}
