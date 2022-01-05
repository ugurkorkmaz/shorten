package routes

import (
	"ugurkorkmaz/shorten/database"

	"github.com/gofiber/fiber/v2"
)

func Resolve(c *fiber.Ctx) error {
	id := c.Params("id")
	bag := database.BAG
	ctx := database.CTX

	url, err := bag.Get(ctx, id).Result()
	if err != nil {
		return c.Status(500).SendString("No Shorten Found with ID")
	}

	return c.Redirect(url)

}
