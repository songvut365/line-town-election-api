package handler

import (
	"bytes"
	"fmt"
	"line-town-election-api/database"
	"line-town-election-api/middleware"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("../template", ".html"),
	})

	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.SetupDatabase()

	return app
}

func TestCreateCandidate(t *testing.T) {
	app := SetUp()
	app.Post("/api/candidates", middleware.Protected, CreateCandidate)

	rand.Seed(time.Now().UnixNano())
	random := fmt.Sprintf("%d", rand.Intn(100))

	body := `{
		"name": "Brown ` + random + `",
		"dob": "August 8, 2011",
		"bioLink": "https://line.fandom.com/wiki/Brown",
		"imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
		"policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown"
		}`

	bodyConvert := []byte(body)

	request, _ := http.NewRequest("POST", "/api/candidates/", bytes.NewBuffer(bodyConvert))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Create Candidate Route")
	assert.Nilf(t, err, "Create Candidate Route")
}

func TestGetCandidates(t *testing.T) {
	app := SetUp()
	app.Get("/api/candidates", middleware.Protected, GetCandidates)

	request, _ := http.NewRequest("GET", "/api/candidates", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Candidates Route")
	assert.Nilf(t, err, "Get Candidates Route")
}

func TestGetCandidate(t *testing.T) {
	app := SetUp()
	app.Get("/api/candidates/:candidateId", middleware.Protected, GetCandidates)

	request, _ := http.NewRequest("GET", "/api/candidates/1", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Candidate Route")
	assert.Nilf(t, err, "Get Candidate Route")
}

func TestUpdateCandidate(t *testing.T) {
	app := SetUp()
	app.Put("/api/candidates/:candidateId", middleware.Protected, UpdateCandidate)

	rand.Seed(time.Now().UnixNano())
	random := fmt.Sprintf("%d", rand.Intn(100))

	body := `{
		"id": 1,
		"name": "Brown ` + random + `",
		"dob": "August 8, 2011",
		"bioLink": "https://line.fandom.com/wiki/Brown",
		"imageLink": "https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630",
		"policy": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown"
		}`

	bodyConvert := []byte(body)

	request, _ := http.NewRequest("PUT", "/api/candidates/1", bytes.NewBuffer(bodyConvert))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Update Candidate Route")
	assert.Nilf(t, err, "Update Candidate Route")
}

func TestDeleteCandidate(t *testing.T) {
	app := SetUp()
	app.Delete("/api/candidates/:candidateId", middleware.Protected, DeleteCandidate)

	request, _ := http.NewRequest("DELETE", "/api/candidates/1", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Delete Candidate Route")
	assert.Nilf(t, err, "Delete Candidate Route")
}
