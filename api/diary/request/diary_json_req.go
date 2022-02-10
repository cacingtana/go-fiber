package request

import (
	"go-fiber/app/models"
	"time"
)

type DiaryRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (request *DiaryRequest) CreateDiary() *models.Diary {
	var diaryRequest models.Diary
	diaryRequest.Title = request.Title
	diaryRequest.Body = request.Body
	diaryRequest.CreateAt = time.Now().UTC()
	return &diaryRequest
}

func (request *DiaryRequest) UpdateDiary() *models.Diary {
	var diaryRequest models.Diary
	diaryRequest.Title = request.Title
	diaryRequest.Body = request.Body
	diaryRequest.CreateAt = time.Now().UTC()
	return &diaryRequest
}
