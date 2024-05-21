package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(404).JSON(response.Web{
		Message: "NOT FOUND",
		Data:    "The requested resource could not be found.",
	})
}
