package user

import (
	"fmt"
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

func (s *servive) Register(email string, password string, level int) error {
	pass := s.hashing.MD5Hash(password)
	data := models.User{
		Email:    email,
		Password: pass,
		Level:    level,
		CreateAt: time.Now().UTC(),
	}

	fmt.Println(&data)
	err := s.repository.Register(data)
	if err != nil {
		return err
	}

	return nil
}
