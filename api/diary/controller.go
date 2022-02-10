package diary

import (
	"go-fiber/api/diary/request"
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
	diary, err := c.service.GetDiaryById(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(diary)
}

func (c *Controller) GetAllDiary(f *fiber.Ctx) error {
	diary, err := c.service.GetAllDiary()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(diary)
}

func (c *Controller) CreateDiary(f *fiber.Ctx) error {
	diary := new(request.DiaryRequest)
	if err := f.BodyParser(&diary); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := c.service.CreateDiary(*diary.CreateDiary()); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success added diary",
	})

}

func (c *Controller) UpdateDiary(f *fiber.Ctx) error {
	diary := new(request.DiaryRequest)
	data := f.Params("id")
	id, _ := strconv.Atoi(data)
	if err := f.BodyParser(&diary); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := c.service.UpdateDiary(*diary.UpdateDiary(), id); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success update diary",
	})
}

func (c *Controller) DeleteDiary(f *fiber.Ctx) error {
	data := f.Params("id")
	id, _ := strconv.Atoi(data)

	if err := c.service.DeleteDiary(id); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success delete diary",
	})
}
