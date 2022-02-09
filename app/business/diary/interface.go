package diary

import (
	"go-fiber/app/models"
)

type Service interface {
	GetDiaryById(id int) (*models.Diary, error)
	GetAllDiary() ([]models.Diary, error)
	CreateDiary(diary models.Diary) error
	UpdateDiary(diary models.Diary, id int) error
	DeleteDiary(id int) error
}

type Repository interface {
	GetDiaryById(id int) (*models.Diary, error)
	GetAllDiary() ([]models.Diary, error)
	CreateDiary(diary models.Diary) error
	UpdateDiary(diary models.Diary, id int) error
	DeleteDiary(id int) error
}
