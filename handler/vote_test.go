package handler

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckVouteStatus(t *testing.T) {
	app := SetUp()
	app.Post("/api/vote/status", CheckVouteStatus)

	body := `{
		"nationalId": "99900023123124"
		}`

	bodyConvert := []byte(body)

	request, _ := http.NewRequest("POST", "/api/vote/status", bytes.NewBuffer(bodyConvert))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Check Vote Status Route")
	assert.Nilf(t, err, "Check Vote Status Route")
}

func TestVote(t *testing.T) {
	app := SetUp()
	app.Post("/api/vote", Vote)

	body := `{
		"nationalId": "123124566123",
		"candidateId": 1
		}`

	bodyConvert := []byte(body)

	request, _ := http.NewRequest("POST", "/api/vote", bytes.NewBuffer(bodyConvert))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer xxxx")

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Vote Route")
	assert.Nilf(t, err, "Vote Route")
}
