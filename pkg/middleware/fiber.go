package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
	)
}
