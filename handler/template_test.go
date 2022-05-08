package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopTenChart(t *testing.T) {
	app := SetUp()
	app.Get("/", GetTopTenChart)

	request, _ := http.NewRequest("GET", "/", nil)

	response, err := app.Test(request, -1)
	assert.Equalf(t, http.StatusOK, response.StatusCode, "Get Top Ten Chart Route")
	assert.Nilf(t, err, "Get Top Ten Chart Route")
}
