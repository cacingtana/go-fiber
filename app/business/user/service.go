package user

import (
	"go-fiber/app/models"
	"go-fiber/pkg/utils"
	"time"
)

type servive struct {
	repository Repository
	hashing    utils.Hash
}

func NewService(repository Repository) Service {
	return &servive{
		repository,
		utils.Hash{},
	}
}

func (s *servive) Login(email string, password string) (*models.User, error) {
	pass := s.hashing.MD5Hash(password)
	login, err := s.repository.Login(email, pass)
	if err != nil {
		return nil, err
	}
	return login, err
}

func (s *servive) Register(user models.User) (*models.UserCredential, error) {
	pass := s.hashing.MD5Hash(user.Password)
	data := &models.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: pass,
		CreateAt: time.Now().UTC(),
	}
	result, err := s.repository.Register(*data)
	if err != nil {
		return nil, err
	}

	return &models.UserCredential{
		Email:  result.Email,
		Status: result.Status,
	}, nil
}
