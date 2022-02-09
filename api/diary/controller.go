package diary

import (
	"go-fiber/app/business/diary"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service diary.Service
}

func NewController(service diary.Service) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) GetDiaryById(f *fiber.Ctx) error {
	data := f.Params("id")
	id, _ := strconv.Atoi(data)
	_, err := c.service.GetDiaryById(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
