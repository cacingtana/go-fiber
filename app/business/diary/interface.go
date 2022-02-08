package diary

import (
	"go-fiber/app/models"
)

type Service interface {
	GetDiaryById(id uint) (*models.Diary, error)
	GetAllDiary() ([]models.Diary, error)
	CreateDiary(diary models.Diary) error
	UpdateDiary(diary models.Diary, id uint) error
	DeleteDiary(id uint) error
}

type Repository interface {
	GetDiaryById(id uint) (*models.Diary, error)
	GetAllDiary() ([]models.Diary, error)
	CreateDiary(diary models.Diary) error
	UpdateDiary(diary models.Diary, id uint) error
	DeleteDiary(id uint) error
}
