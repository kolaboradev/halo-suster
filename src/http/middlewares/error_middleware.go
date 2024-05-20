package middlewares

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
)

func ErrorHandle(c *fiber.Ctx, err error) error {
	if value, ok := err.(*exceptions.ConflictError); ok {
		valErr := value.Error()
		c.Status(409)
		return c.JSON(response.Web{
			Message: "CONFLICT ERROR",
			Data:    valErr,
		})
	}
	if value, ok := err.(*exceptions.BadRequestError); ok {
		valErr := value.Error()
		c.Status(400)
		return c.JSON(response.Web{
			Message: "BAD REQUEST",
			Data:    valErr,
		})
	}
	if value, ok := err.(*exceptions.NotFoundError); ok {
		valErr := value.Error()
		c.Status(404)
		return c.JSON(response.Web{
			Message: "NOT FOUND",
			Data:    valErr,
		})
	}
	if value, ok := err.(validator.ValidationErrors); ok {
		valErr := value[0]
		message := helper.CustomMessageValidation(valErr)
		c.Status(400)
		return c.JSON(response.Web{
			Message: "Validation Error",
			Data:    message,
		})
	}

	fmt.Println(err)
	c.Status(500)
	return c.JSON(response.Web{
		Message: "SERVER ERROR",
		Data:    err.Error(),
	})
}
