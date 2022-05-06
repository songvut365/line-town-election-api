package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// POST Check Vote status
// API to check vote status for voter
func CheckVouteStatus(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
	})
}

// POST Vote
// API to vote
func Vote(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
