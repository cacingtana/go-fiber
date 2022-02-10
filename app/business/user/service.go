package user

import (
	"go-fiber/api/user/response"
	"go-fiber/app/models"
	"go-fiber/pkg/middleware"
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

func (s *servive) Login(email string, password string) (*response.LoginResponse, error) {
	pass := s.hashing.MD5Hash(password)
	data, err := s.repository.Login(email, pass)
	if err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(data.Level, data.Email)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Email: data.Email,
		Token: token,
	}, nil
}

func (s *servive) Register(email string, password string, level int) error {
	pass := s.hashing.MD5Hash(password)
	data := models.User{
		Email:    email,
		Password: pass,
		Level:    level,
		CreateAt: time.Now().UTC(),
	}
	err := s.repository.Register(data)
	if err != nil {
		return err
	}

	return nil
}
