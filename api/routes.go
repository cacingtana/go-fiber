package api

import (
	"go-fiber/api/diary"
	"go-fiber/api/user"
	"go-fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Router(f *fiber.App, userController *user.Controller, diaryController *diary.Controller) {

	//use middleware
	route := f.Group("/api")
	route.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	route.Use(middleware.JWTAuth())
	route.Get("/diary", diaryController.GetAllDiary)
	route.Get("/diary/:id", diaryController.GetDiaryById)
	route.Post("/diary", diaryController.CreateDiary)
	route.Put("/diary/:id", diaryController.UpdateDiary)
	route.Delete("diary/:id", diaryController.DeleteDiary)

	//public route
	f.Post("/login", userController.Login)
	f.Post("/register", userController.Register)

}
