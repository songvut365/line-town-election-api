package handler

import (
	"line-town-election-api/database"
	"line-town-election-api/model"
	"line-town-election-api/validation"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// POST Check Vote status
// API to check vote status for voter
func CheckVouteStatus(c *fiber.Ctx) error {
	db := database.Database

	// Parser
	var input model.InputCheckVote

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
		})
	}

	// Validation
	errors := validation.ValidInput(input)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"error":   errors,
		})
	}

	// Find vote by nationalId
	var vote model.Vote

	result := db.Where("national_id = ?", input.NationalID).First(&vote)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": false,
		})
	} else {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": true,
		})
	}

}

// POST Vote
// API to vote
func Vote(c *fiber.Ctx) error {
	db := database.Database

	// Parser
	var input model.InputVote

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
		})
	}

	// Validation
	errors := validation.ValidInput(input)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"error":   errors,
		})
	}

	// Check election is closed
	if !ElectionStatus {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Election is closed",
		})
	}

	// Check already voted
	var vote model.Vote

	err = db.Where("national_id = ?", input.NationalID).First(&vote).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Already voted",
		})
	}

	// Find candidate
	var candidate model.Candidate

	err = db.Where("id = ?", input.CandidateID).First(&candidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Candidate not found",
		})
	}

	// Vote
	var newVote = model.Vote{
		NationalID:  input.NationalID,
		CandidateID: input.CandidateID,
	}

	err = db.Create(&newVote).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot vote",
		})
	}

	// Count vote
	var voteCount int64

	db.Model(&model.Vote{}).Where("candidate_id = ?", input.CandidateID).Count(&voteCount)

	// Update vote count
	err = db.Model(&model.Candidate{}).Where("id = ?", input.CandidateID).Update("voted_count", uint(voteCount)).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot update vote count",
		})
	}

	//
	var logVote model.LogVote
	logVote.ID = candidate.ID
	logVote.VotedCount = uint(voteCount)
	logVote.CreatedAt = time.Now()

	db.Create(&logVote)

	// Success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
