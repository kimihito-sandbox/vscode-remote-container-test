package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})

	app.Listen(3000)
}
