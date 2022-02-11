package diary

import (
	"go-fiber/app/business"
	MockDiary "go-fiber/app/business/diary/mock"
	"go-fiber/app/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	id       = 1
	title    = "Golang"
	body     = "Belajar pemrograman golang"
	createAt = true
)

var (
	diaryServ Service
	diaryRepo MockDiary.Repository

	diarysData []models.Diary = make([]models.Diary, 0)

	diaryData   models.Diary
	createDiary models.Diary
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	createDiary = models.Diary{
		Id:    id,
		Title: title,
		Body:  body,
	}

	diaryServ = NewService(&diaryRepo)

}

func TestGetDiaryById(t *testing.T) {
	t.Run("Expect found the diary", func(t *testing.T) {
		diaryRepo.On("GetDiaryById", mock.AnythingOfType("int")).Return(&diaryData, nil).Once()

		diary, err := diaryServ.GetDiaryById(id)

		assert.Nil(t, err)
		assert.NotNil(t, diary)
	})

	t.Run("Expect diary not found", func(t *testing.T) {
		diaryRepo.On("GetDiaryById", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		product, err := diaryServ.GetDiaryById(id)

		assert.NotNil(t, err)
		assert.Nil(t, product)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestGetAllDiary(t *testing.T) {
	t.Run("Expect found all diary", func(t *testing.T) {
		diaryRepo.On("GetAllDiary", mock.Anything).Return(diarysData, nil).Once()

		diarys, err := diaryServ.GetAllDiary()

		assert.Nil(t, err)
		assert.NotNil(t, diarys)

	})
}

func TestCreateDiary(t *testing.T) {
	t.Run("Expect create diary success", func(t *testing.T) {
		diaryRepo.On("CreateDiary", mock.AnythingOfType("models.Diary"), mock.AnythingOfType("string")).Return(nil).Once()

		err := diaryServ.CreateDiary(createDiary)
		assert.Nil(t, err)

	})

	t.Run("Expect create diary not found", func(t *testing.T) {
		diaryRepo.On("CreateDiary", mock.AnythingOfType("models.Diary"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := diaryServ.CreateDiary(createDiary)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestUpdateDiary(t *testing.T) {

}
func TestDeleteDiary(t *testing.T) {
	t.Run("Success delete diary", func(t *testing.T) {
		diaryRepo.On("DeleteDiary", mock.AnythingOfType("int")).Return(nil).Once()
		err := diaryServ.DeleteDiary(id)
		assert.Nil(t, err)
	})
}
