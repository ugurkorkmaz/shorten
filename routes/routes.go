package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/shorten", Shorten)
	app.Get("/:id", Resolve)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}
