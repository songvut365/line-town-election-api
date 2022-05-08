package handler

import (
	"bytes"
	"line-town-election-api/middleware"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToggleElection(t *testing.T) {
	app := SetUp()
	app.Post("/api/election/toggle", middleware.Protected, ToggleElection)

	body := `{
		"enable": true
		}`

	bodyConvert := []byte(body)

	request, _ := http.NewRequest("POST", "/api/election/toggle", bytes.NewBuffer(bodyConvert))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Toggle Election Route")
	assert.Nilf(t, err, "Toggle Election Route")
}

func TestGetElectionCount(t *testing.T) {
	app := SetUp()
	app.Get("/api/election/count", middleware.Protected, GetElectionCount)

	request, _ := http.NewRequest("GET", "/api/election/count", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Election Count Route")
	assert.Nilf(t, err, "Get Election Count Route")
}

func TestGetElectionResult(t *testing.T) {
	app := SetUp()
	app.Get("/api/election/result", middleware.Protected, GetElectionResult)

	request, _ := http.NewRequest("GET", "/api/election/result", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Election Result Route")
	assert.Nilf(t, err, "Get Election Result Route")
}

func TestGetExportResult(t *testing.T) {
	app := SetUp()
	app.Get("/api/election/export", middleware.Protected, GetExportResult)

	request, _ := http.NewRequest("GET", "/api/election/export", nil)
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Export Result Route")
	assert.Nilf(t, err, "Get Export Result Route")
}
