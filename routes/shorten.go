package routes

import (
	"log"
	"ugurkorkmaz/shorten/database"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

/*var (
	validate *validator.Validate
)*/

func Shorten(c *fiber.Ctx) error {
	url := c.Query("url")
	/*url = reflect.TypeOf(url).String()
	err := validate.Var(url, "required,url")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "url is required and must be a valid url",
			"code":  err,
		})
	}*/
	uid := uuid.New().String()[:6]

	dbs := database.DB
	dbs.AutoMigrate(&database.Shorten{})
	dbs.Create(&database.Shorten{
		Long: url,
		Uid:  uid,
	})
	bag := database.BAG
	ctx := database.CTX

	err := bag.Set(ctx, uid, url, 0).Err()
	if err != nil {
		panic(err)
	}
	shorten := database.Shorten{}
	dbs.First(&shorten, "uid = ?", uid)
	log.Print(shorten)
	return c.JSON(shorten)
}
