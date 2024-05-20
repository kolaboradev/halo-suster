package helper

import (
	"github.com/gofiber/fiber/v2"
)

func QueryIntPointer(c *fiber.Ctx, key string) *int {
	value := c.QueryInt(key, -1)
	if value != -1 {
		return &value
	}
	return nil
}

func QueryInt64Pointer(c *fiber.Ctx, key string) *int64 {
	value := c.QueryInt(key, -1)
	if value != -1 {
		value64 := int64(value)
		return &value64
	}
	return nil
}
