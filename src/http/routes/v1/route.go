package v1Routes

import "github.com/gofiber/fiber/v2"

func SetRoutesV1(app *fiber.App) fiber.Router {
	v1 := app.Group("/v1")
	return v1
}
