package recordController

import "github.com/gofiber/fiber/v2"

type RecordControllerInterface interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}
