package user

import "go-fiber/app/models"

type Service interface {
	Login(email string, password string) (*models.User, error)
	Register(email string, password string) (*models.UserCredential, error)
}

type Repository interface {
	Login(email string, password string) (*models.User, error)
	Register(email string, password string) (*models.UserCredential, error)
}
