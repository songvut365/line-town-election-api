package handler

import (
	"line-town-election-api/database"
	"line-town-election-api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

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

func GetCandidate(c *fiber.Ctx) error {
	db := database.Database

	// Find candidate by id
	candidateId := c.Params("candidateId")

	var candidate model.Candidate

	err := db.Where("id = ?", candidateId).First(&candidate).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get candidate",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(candidate)
}

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

	// Find exist candidate by name
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot update candidate",
		})
	}

	return c.Status(http.StatusOK).JSON(candidate)
}

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
