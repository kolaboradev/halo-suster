package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kolaboradev/halo-suster/src/helper"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Web{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}
	token, err := jwt.Parse(tokenStr, helper.CheckTokenJWT)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Web{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Web{
			Message: "UNAUTHORIZED",
			Data:    "Token Invalid",
		})
	}

	c.Locals("userId", token.Claims.(jwt.MapClaims)["userId"])
	c.Locals("nip", token.Claims.(jwt.MapClaims)["nip"])
	c.Locals("role", token.Claims.(jwt.MapClaims)["role"])

	return c.Next()
}
