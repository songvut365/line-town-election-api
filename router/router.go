package router

import (
	"line-town-election-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app fiber.App) {
	api := app.Group("/api")

	//Candidate Routes
	api.Get("/candidates", handler.GetCandidates)
	api.Get("/candidates/:candidateId", handler.GetCandidate)
	api.Post("/candidates", handler.CreateCandidate)
	api.Put("/candidates/:candidateId", handler.UpdateCandidate)
	api.Delete("/candidates/:candidateId", handler.DeleteCandidate)

	//Vote Routes
	api.Post("/vote/status", handler.CheckVouteStatus)
	api.Post("/vote", handler.Vote)

	//Election Routes
	api.Post("/election/toggle", handler.ToggleElection)
	api.Get("/election/count", handler.GetElectionCount)
	api.Get("/election/result", handler.GetElectionResult)
	api.Get("/election/export", handler.GetExportResult)
}
