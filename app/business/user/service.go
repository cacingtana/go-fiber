package user

import (
	"go-fiber/app/models"
	"go-fiber/pkg/utils"
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

func (s *servive) Register(email string, password string) (*models.UserCredential, error) {
	pass := s.hashing.MD5Hash(password)
	return s.repository.Register(email, pass)
}
