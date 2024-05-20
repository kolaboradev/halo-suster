package userController

import "github.com/gofiber/fiber/v2"

type UserControllerInterface interface {
	RegisterIt(c *fiber.Ctx) error
	LoginIt(c *fiber.Ctx) error
	RegisterNurse(c *fiber.Ctx) error
	LoginNurse(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	EditNurseById(c *fiber.Ctx) error
	DeleteNurseById(c *fiber.Ctx) error
	GetAccessNurse(c *fiber.Ctx) error
	PostImage(c *fiber.Ctx) error
}
