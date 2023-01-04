package routes

import (
	"github.com/Nameer-kp/go-url-shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.GetRDBClient0()
	value, err := r.Get(c.Context(), url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the database"})

	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to db"})
	}
	rInr := database.GetRDBClient1()
	_ = rInr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}
