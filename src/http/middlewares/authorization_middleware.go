package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
)

func AuthorizeMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		fmt.Println(role)
		for _, value := range roles {
			if role == value {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(response.Web{
			Message: "FORBIDDEN",
			Data:    "You do not have the required permissions",
		})
	}
}
