package user

import (
	"go-fiber/app/models"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserQuery {
	return &UserQuery{
		db,
	}
}

func (q *UserQuery) Login(email string, password string) (*models.User, error) {
	var data models.User
	if err := q.db.Where("email = ?", email).Where("password = ?", password).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *UserQuery) Register(user *models.User) (*models.UserCredential, error) {
	if err := q.db.Save(user).Error; err != nil {
		return nil, err
	}
	data := &models.UserCredential{
		Email:  user.Email,
		Status: "User Registered",
	}
	return data, nil
}
