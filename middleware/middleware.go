package middleware

import (
	"github.com/ernestokorpys/gobackend/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unautheticated",
		})
	}
	return c.Next()
}
