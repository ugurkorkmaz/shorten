package main

import (
	"os"
	"ugurkorkmaz/shorten/config"
	"ugurkorkmaz/shorten/database"
	"ugurkorkmaz/shorten/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	config.Update()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
