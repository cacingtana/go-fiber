package diary

import (
	"go-fiber/app/models"

	"gorm.io/gorm"
)

type DiaryQuery struct {
	db *gorm.DB
}

func NewDiaryRepo(db *gorm.DB) *DiaryQuery {
	return &DiaryQuery{
		db,
	}
}

func (q *DiaryQuery) GetDiaryById(id int) (*models.Diary, error) {
	var data models.Diary
	if err := q.db.First(&data, id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *DiaryQuery) GetAllDiary() ([]models.Diary, error) {
	var data, result []models.Diary
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	for _, v := range data {
		result = append(result, v)
	}
	return result, nil
}

func (q *DiaryQuery) CreateDiary(diary models.Diary) error {
	if err := q.db.Save(diary).Error; err != nil {
		return err
	}
	return nil
}

func (q *DiaryQuery) UpdateDiary(diary models.Diary, id int) error {
	if err := q.db.Where("id = ?", &id).Updates(diary).Error; err != nil {
		return err
	}
	return nil
}
func (q *DiaryQuery) DeleteDiary(id int) error {
	var data models.Diary
	if err := q.db.Delete(&data, id).Error; err != nil {
		return err
	}
	return nil
}
