package main

import (
	"go-fiber/configs"
	"go-fiber/pkg/middleware"
	"log"

	"github.com/gofiber/fiber/v2"

	diaryController "go-fiber/api/diary"
	diaryService "go-fiber/app/business/diary"
	diaryRepository "go-fiber/app/repository/diary"

	userController "go-fiber/api/user"
	userService "go-fiber/app/business/user"
	userRepository "go-fiber/app/repository/user"

	f "go-fiber/api"
)

func main() {

	connection := configs.MySQLConn()

	diaryRepo := diaryRepository.NewDiaryRepo(connection)
	diaryServ := diaryService.NewService(diaryRepo)
	diaryCont := diaryController.NewController(diaryServ)

	userRepo := userRepository.NewUserRepo(connection)
	userServ := userService.NewService(userRepo)
	userCont := userController.NewController(userServ)

	config := configs.ServerTimeOut()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)
	f.Router(app, userCont, diaryCont)

	log.Fatal(app.Listen(":8000"))
}
