package request

import "go-fiber/app/models"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    int    `json:"password"`
}

func (request *RegisterRequest) Register() *models.User {
	var userRequest models.User
	userRequest.Email = request.Email
	userRequest.Password = request.Password
	userRequest.Level = request.Level
	return &userRequest
}
