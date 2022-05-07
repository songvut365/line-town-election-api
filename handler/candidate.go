package handler

import (
	"line-town-election-api/database"
	"line-town-election-api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET Candidate
// API to get Candidate list
func GetCandidates(c *fiber.Ctx) error {
	db := database.Database

	// Find candidates
	candidates := []model.Candidate{}

	err := db.Find(&candidates).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get candidates",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(candidates)
}

// GET Candidate Detail
// API to get Candidate detail
func GetCandidate(c *fiber.Ctx) error {
	db := database.Database

	// Find candidate by id
	candidateId := c.Params("candidateId")

	var candidate model.Candidate

	err := db.Where("id = ?", candidateId).First(&candidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Candidate not found",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(candidate)
}

// POST Create a new Candidate
// API to create a new Candidate
func CreateCandidate(c *fiber.Ctx) error {
	db := database.Database

	// Parser
	var candidate model.InputCandidate

	err := c.BodyParser(&candidate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parser body",
		})
	}

	// Find exist candidate by name
	var existCandidate model.Candidate

	err = db.Where("name = ?", candidate.Name).First(&existCandidate).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Candidate name already exist",
		})
	}

	// Create Candidate
	var newCandidate = model.Candidate{
		Name:      candidate.Name,
		DOB:       candidate.DOB,
		BioLink:   candidate.BioLink,
		ImageLink: candidate.ImageLink,
		Policy:    candidate.Policy,
	}

	err = db.Model(&model.Candidate{}).Create(&newCandidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create new candidate",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(newCandidate)
}

// PUT Update a Candidate
// API to update a Candidate
func UpdateCandidate(c *fiber.Ctx) error {
	db := database.Database

	// Parser
	candidateId := c.Params("candidateId")

	var input model.Candidate

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parser body",
		})
	}

	// Voted coutn can't update
	if input.VotedCount != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Voted count cannot update",
		})
	}

	// Find exist candidate by id
	var candidate model.Candidate

	err = db.Where("id = ?", candidateId).First(&candidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Candidate does not exist",
		})
	}

	// Update Candidate
	err = db.Model(&candidate).Updates(&input).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot update candidate",
		})
	}

	// Update candidate id of votes
	err = db.Model(&model.Vote{}).Where("candidate_id = ?", candidateId).Update("candidate_id", candidate.ID).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot update candidate id in votes",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(candidate)
}

// DELETE Delete a Candidate
// API to delete a Candidate
func DeleteCandidate(c *fiber.Ctx) error {
	db := database.Database

	// Find exist candidate
	candidateId := c.Params("candidateId")

	var candidate model.Candidate

	err := db.Where("id = ?", candidateId).First(&candidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Candidate not found",
		})
	}

	// Delete candidate by id
	err = db.Delete(&candidate).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot delete candidate",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
