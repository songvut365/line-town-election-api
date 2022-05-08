package handler

import (
	"line-town-election-api/database"
	"line-town-election-api/model"

	"github.com/gofiber/fiber/v2"
)

func GetTopTenChart(c *fiber.Ctx) error {
	colors := []string{
		"#4CC764", "#f1c40f", "#3498db", "#e74c3c", "#9b59b6",
		"#2ecc71", "#34495e", "#bdc3c7", "#7f8c8d", "#f39c12",
	}

	db := database.Database
	candidates := []model.Candidate{}
	db.Limit(10).Find(&candidates)

	data := []uint{}
	label := []string{}
	color := []string{}

	for index, candidate := range candidates {
		data = append(data, *candidate.VotedCount)
		label = append(label, candidate.Name)
		color = append(color, colors[index])
	}

	return c.Render("index", fiber.Map{
		"Data":  data,
		"Label": label,
		"Color": color,
	})
}
