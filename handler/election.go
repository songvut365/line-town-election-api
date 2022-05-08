package handler

import (
	"encoding/csv"
	"fmt"
	"line-town-election-api/database"
	"line-town-election-api/model"
	"line-town-election-api/validation"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var ElectionStatus bool // Election Status injected from main

// POST Toggle Election
// API to open or close election
func ToggleElection(c *fiber.Ctx) error {
	// Parser
	var input model.InputToggleElection

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

	// Toggle
	ElectionStatus = *input.Enable

	// Success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"enable": ElectionStatus,
	})
}

// POST Election Count
// API to get id and voted count of candidate
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

// GET Election Result
// API to get candidate and persentage
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
	var percentage string
	electionResults := []model.ResponseElectionResult{}

	for _, candidate := range candidates {

		if votedAll == 0 {
			percentage = "0%"
		} else {
			total := float64(*candidate.VotedCount) / float64(votedAll)
			percentage = fmt.Sprintf("%.2f", total*100) + "%" //convert to string
		}

		result := model.ResponseElectionResult{
			ID:         candidate.ID,
			Name:       candidate.Name,
			DOB:        candidate.DOB,
			BioLink:    candidate.BioLink,
			ImageLink:  candidate.ImageLink,
			Policy:     candidate.Policy,
			VotedCount: *candidate.VotedCount,
			Percentage: percentage,
		}

		electionResults = append(electionResults, result)
	}

	// Success
	return c.Status(http.StatusOK).JSON(electionResults)
}

// GET Exported Result (download)
// API to send csv file with national id and candidate id
func GetExportResult(c *fiber.Ctx) error {
	db := database.Database

	// Get all votes
	var votes = []model.Vote{}

	db.Find(&votes)

	// Create csv file
	file, err := os.Create(os.Getenv("CSV_FILE"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot open csv file",
		})
	}
	defer file.Close()

	// Create csv writer and title
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string

	title := []string{"Candidate id", "National id"}
	data = append(data, title)

	// Write votes to csv file
	if len(votes) != 0 {
		for _, vote := range votes {
			candidateId := fmt.Sprintf("%v", vote.CandidateID)
			nationalId := vote.NationalID

			row := []string{candidateId, nationalId}
			data = append(data, row)
		}
	}

	writer.WriteAll(data)

	// Success
	return c.Download(os.Getenv("CSV_FILE_SEND"), "result.csv")
}

// Real-time Vote Stream
// Websocket stream for real-time vote count
func CandidateVoteStream(ws *websocket.Conn) {
	db := database.Database

	candidateId := ws.Params("candidateId")

	// Variable for before query candidate
	var before model.ResponseElectionCount

	for {
		var now model.ResponseElectionCount

		db.Model(&model.LogVote{}).Select("id", "voted_count").Where("id = ?", candidateId).Order("created_at desc").First(&now)

		// Check voted count has changed
		if now.VotedCount != before.VotedCount {
			err := ws.WriteJSON(now)
			if err != nil {
				break
			}

			before = now
		}
	}
}

// Real-time Vote Stream
// Websocket stream for real-time vote count
func CandidatesVoteStream(ws *websocket.Conn) {
	db := database.Database

	// Variable for before query candidate
	var before model.ResponseElectionCount

	for {
		var now model.ResponseElectionCount

		db.Model(&model.LogVote{}).Select("id", "voted_count").Order("created_at desc").First(&now)

		// Check voted count has changed
		if now.VotedCount != before.VotedCount {
			err := ws.WriteJSON(now)
			if err != nil {
				break
			}

			before = now
		}
	}
}
