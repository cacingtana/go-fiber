package user

import (
	"fmt"
	"go-fiber/api/user/request"
	"go-fiber/app/business/user"
	"reflect"

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
	login := new(request.LoginRequest)
	if err := f.BodyParser(login); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	user, err := c.service.Login(login.Email, login.Password)

	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(user)
}

func (c *Controller) Register(f *fiber.Ctx) error {
	register := new(request.RegisterRequest)
	if err := f.BodyParser(&register); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fmt.Println(reflect.TypeOf(register))
	err := c.service.Register(register.Email, register.Password, register.Level)

	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "User Registered",
	})
}
