package middleware

import (
	"go_crud/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {

	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorization",
		})
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorization",
		})
	}

	// if token != "humam" {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unauthorization",
	// 	})
	// }


	return ctx.Next()
}

func GenerateToken(ctx *fiber.Ctx) error {

	return ctx.Next()
}