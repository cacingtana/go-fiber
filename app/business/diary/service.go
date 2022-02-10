package diary

import (
	"fmt"
	"go-fiber/app/models"
)

type servive struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &servive{
		repository,
	}
}

func (s *servive) GetDiaryById(id int) (*models.Diary, error) {
	diaryById, err := s.repository.GetDiaryById(id)
	if err != nil {
		return nil, err
	}
	return diaryById, nil
}

func (s *servive) GetAllDiary() ([]models.Diary, error) {
	diary, err := s.repository.GetAllDiary()
	if err != nil {
		return []models.Diary{}, nil
	}
	fmt.Println(diary)
	return diary, nil
}

func (s *servive) CreateDiary(diary models.Diary) error {
	if err := s.repository.CreateDiary(diary); err != nil {
		return err
	}
	return nil
}

func (s *servive) UpdateDiary(diary models.Diary, id int) error {
	if err := s.repository.UpdateDiary(diary, id); err != nil {
		return err
	}
	return nil
}

func (s *servive) DeleteDiary(id int) error {
	if err := s.repository.DeleteDiary(id); err != nil {
		return err
	}
	return nil
}
