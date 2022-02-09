package request

import (
	"go-fiber/app/models"
	"time"
)

type DiaryResponse struct {
	Id       int       `json:"Id"`
	Title    string    `json:"Title"`
	Body     string    `json:"Body"`
	CreateAt time.Time `json:"createAt"`
}

type AllDiaryResponse struct {
	Diary []DiaryResponse
}

func GetDiaryByID(diary models.Diary) *DiaryResponse {
	var diaryResponse DiaryResponse
	diaryResponse.Title = diary.Title
	diaryResponse.Body = diary.Body
	diaryResponse.CreateAt = diary.CreateAt
	return &diaryResponse
}

func GetAllDiary(diary []models.Diary) *AllDiaryResponse {
	var diaryResponse DiaryResponse

	var allResponse AllDiaryResponse
	for _, val := range diary {
		diaryResponse.Title = val.Title
		diaryResponse.Body = val.Body
		diaryResponse.CreateAt = val.CreateAt
		allResponse.Diary = append(allResponse.Diary, diaryResponse)
	}
	return &allResponse
}
