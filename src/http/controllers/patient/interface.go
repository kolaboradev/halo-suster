package patientController

import "github.com/gofiber/fiber/v2"

type PatientControllerInterface interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}
