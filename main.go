package main

import (
	"go-fiber/configs"
	"go-fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.ServerTimeOut()
	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(":3000")
}
