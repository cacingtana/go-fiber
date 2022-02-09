package user

import (
	"go-fiber/app/business/user"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) Login(f *fiber.Ctx) error {
	return nil
}
