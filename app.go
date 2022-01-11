package main

import (
	"os"
	"ugurkorkmaz/shorten/config"
	"ugurkorkmaz/shorten/database"
	"ugurkorkmaz/shorten/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	database.Connect()
	config.Update()
	engine := html.New("./template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	routes.SetupRoutes(app)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
