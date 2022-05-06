package handler

import (
	"line-town-election-api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// POST Toggle Election
func ToggleElection(c *fiber.Ctx) error {
	//Success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"enable": true,
	})
}

// Get Election Count
func GetElectionCount(c *fiber.Ctx) error {
	electionCounts := []model.ResponseElectionCount{}

	//Success
	return c.Status(http.StatusOK).JSON(electionCounts)
}

// GET Election Result
func GetElectionResult(c *fiber.Ctx) error {
	electionResults := []model.ResponseElectionResult{}

	//Success
	return c.Status(http.StatusOK).JSON(electionResults)
}

// GET Exported Result (download)
func GetExportResult(c *fiber.Ctx) error {
	//Success
	return c.Download("./public/export/result.csv", "result.csv")
}
