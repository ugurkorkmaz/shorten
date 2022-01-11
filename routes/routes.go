package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Static("/", "./public")
	app.Get("/shorten", Shorten)
	app.Get("/:id", Resolve)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("app", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}
