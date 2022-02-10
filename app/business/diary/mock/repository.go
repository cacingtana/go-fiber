package mock

import (
	"go-fiber/app/models"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (t *Repository) GetDiaryById(id int) (*models.Diary, error) {
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

func (t *Repository) GetAllDiary() ([]models.Diary, error) {
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

func (t *Repository) CreateDiary(_t models.Diary) error {
	ret := t.Called(_t)
	var tError error
	if rf, ok := ret.Get(0).(func(models.Diary) error); ok {
		tError = rf(_t)
	} else {
		tError = ret.Error(0)
	}
	return tError
}

func (t *Repository) UpdateDiary(_t models.Diary, id int) error {
	ret := t.Called(_t, id)
	var tError error
	if rf, ok := ret.Get(0).(func(models.Diary, int) error); ok {
		tError = rf(_t, id)
	} else {
		tError = ret.Error(0)
	}
	return tError
}

func (t *Repository) DeleteDiary(id int) error {
	ret := t.Called(id)
	var tError error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		tError = rf(id)
	} else {
		tError = ret.Error(0)
	}

	return tError
}
