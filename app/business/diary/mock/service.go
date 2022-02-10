package mock

import (
	"go-fiber/app/models"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (t *Service) GetAllDiary() ([]models.Diary, error) {
	ret := t.Called()
	var tSuccess []models.Diary

	if rf, ok := ret.Get(0).(func() []models.Diary); ok {
		tSuccess = rf()
	} else {
		if ret.Get(0) != nil {
			tSuccess = ret.Get(0).([]models.Diary)
		}
	}
	return tSuccess, nil
}

func (t *Service) GetDiaryById(id int) (*models.Diary, error) {
	ret := t.Called(id)

	var tSuccess *models.Diary
	if rf, ok := ret.Get(0).(func(int) *models.Diary); ok {
		tSuccess = rf(id)
	} else {
		if ret.Get(0) != nil {
			tSuccess = ret.Get(0).(*models.Diary)
		}
	}

	var tError error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		tError = rf(id)
	} else {
		tError = ret.Error(1)
	}
	return tSuccess, tError
}

func (t *Service) CreateDiary(diary models.Diary) error {
	ret := t.Called(diary)
	var tError error
	if rf, ok := ret.Get(0).(func(models.Diary) error); ok {
		tError = rf(diary)
	} else {
		tError = ret.Error(0)
	}
	return tError
}

func (t *Service) DeleteDiary(id int) error {
	ret := t.Called(id)
	var tError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tError = rf(id)
	} else {
		tError = ret.Error(0)
	}
	return tError
}

func (t *Service) UpdateDiary(diary models.Diary, id int) error {
	ret := t.Called(diary, id)

	var tError error
	if rf, ok := ret.Get(0).(func(models.Diary, int) error); ok {
		tError = rf(diary, id)
	} else {
		tError = ret.Error(0)
	}

	return tError
}
