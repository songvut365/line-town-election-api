package middleware

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func compareToken(authorization string) bool {
	bearer := "Bearer "
	return authorization == bearer+os.Getenv("EASY_TOKEN")
}

func Protected(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	// Verify Token
	if !compareToken(token) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid token",
		})
	}

	return c.Next()
}