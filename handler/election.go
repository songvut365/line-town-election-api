package handler

import (
	"encoding/csv"
	"fmt"
	"line-town-election-api/database"
	"line-town-election-api/model"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

var ElectionStatus bool // Election Status injected from main

func ToggleElection(c *fiber.Ctx) error {
	// Parser
	var input model.InputToggleElection

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parser body",
		})
	}

	// Toggle
	ElectionStatus = input.Enable

	// Success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"enable": ElectionStatus,
	})
}

func GetElectionCount(c *fiber.Ctx) error {
	db := database.Database

	// Find id and vouted count
	electionCounts := []model.ResponseElectionCount{}

	err := db.Model(&model.Candidate{}).Select("id", "voted_count").Find(&electionCounts).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot find election count",
		})
	}

	// Success
	return c.Status(http.StatusOK).JSON(electionCounts)
}

func GetElectionResult(c *fiber.Ctx) error {
	db := database.Database

	// Find all candidate
	var candidates []model.Candidate

	err := db.Find(&candidates).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot find candidates",
		})
	}

	// Get votes amount
	var votedAll int64

	db.Model(&model.Vote{}).Count(&votedAll)

	// Calculate Percentage
	electionResults := []model.ResponseElectionResult{}

	for _, candidate := range candidates {
		total := float64(candidate.VotedCount) / float64(votedAll)
		percentage := fmt.Sprintf("%.2f", total*100) + "%" //convert to string

		result := model.ResponseElectionResult{
			ID:         candidate.ID,
			Name:       candidate.Name,
			DOB:        candidate.DOB,
			BioLink:    candidate.BioLink,
			ImageLink:  candidate.ImageLink,
			Policy:     candidate.Policy,
			VotedCount: candidate.VotedCount,
			Percentage: percentage,
		}

		electionResults = append(electionResults, result)
	}

	// Success
	return c.Status(http.StatusOK).JSON(electionResults)
}

func GetExportResult(c *fiber.Ctx) error {
	db := database.Database

	// Get all votes
	var votes []model.Vote

	err := db.Find(&votes).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot find votes",
		})
	}

	// Create csv file
	file, err := os.Create("./public/export/result.csv")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot open csv file",
		})
	}
	defer file.Close()

	// Write votes to csv file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string

	title := []string{"Candidate id", "National id"}
	data = append(data, title)

	for _, vote := range votes {
		candidateId := fmt.Sprintf("%v", vote.CandidateID)
		nationalId := vote.NationalID

		row := []string{candidateId, nationalId}
		data = append(data, row)
	}
	writer.WriteAll(data)

	// Success
	return c.Download("./public/export/result.csv", "result.csv")
}
