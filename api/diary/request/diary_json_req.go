package request

import (
	"go-fiber/app/models"
)

type DiaryRequest struct {
	Id    int    `json:"Id"`
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

func (request *DiaryRequest) CreateDiary() *models.Diary {
	var diaryRequest models.Diary
	diaryRequest.Title = request.Title
	diaryRequest.Body = request.Body
	return &diaryRequest
}

func (request *DiaryRequest) UpdateDiary() *models.Diary {
	var diaryRequest models.Diary
	diaryRequest.Id = request.Id
	diaryRequest.Title = request.Title
	diaryRequest.Body = request.Body
	return &diaryRequest
}
