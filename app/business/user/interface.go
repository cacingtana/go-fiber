package user

import (
	"go-fiber/api/user/response"
	"go-fiber/app/models"
)

type Service interface {
	Login(email string, password string) (*response.LoginResponse, error)
	Register(email string, password string, level int) error
}

type Repository interface {
	Login(email string, password string) (*models.User, error)
	Register(user models.User) error
}
