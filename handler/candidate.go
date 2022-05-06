package handler

import (
	"line-town-election-api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET Candidate
// API to get Candidate list
func GetCandidates(c *fiber.Ctx) error {
	candidates := []model.Candidate{}

	return c.Status(http.StatusOK).JSON(candidates)
}

// GET Candidate Detail
// API to get Candidate detail
func GetCandidate(c *fiber.Ctx) error {
	var candidate model.Candidate

	return c.Status(http.StatusOK).JSON(candidate)
}

// POST Create a new Candidate
// API to create a new Candidate
func CreateCandidate(c *fiber.Ctx) error {
	var candidate model.Candidate

	return c.Status(http.StatusOK).JSON(candidate)
}

// PUT Update a Candidate
// API to update a Candidate
func UpdateCandidate(c *fiber.Ctx) error {
	var candidate model.Candidate

	return c.Status(http.StatusOK).JSON(candidate)
}

// DELETE Delete a Candidate
// API to delete a Candidate
func DeleteCandidate(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
