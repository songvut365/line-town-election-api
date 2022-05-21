package router

import (
	"line-town-election-api/handler"
	"line-town-election-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupRouter(app fiber.Router) {
	api := app.Group("/api", middleware.Protected)

	// Candidate Routes
	api.Get("/candidates", handler.GetCandidates)
	api.Get("/candidates/:candidateId", handler.GetCandidate)
	api.Post("/candidates", handler.CreateCandidate)
	api.Put("/candidates/:candidateId", handler.UpdateCandidate)
	api.Delete("/candidates/:candidateId", handler.DeleteCandidate)

	// Vote Routes
	api.Post("/vote/status", handler.CheckVoteStatus)
	api.Post("/vote", handler.Vote)

	// Election Routes
	api.Post("/election/toggle", handler.ToggleElection)
	api.Get("/election/count", handler.GetElectionCount)
	api.Get("/election/result", handler.GetElectionResult)
	api.Get("/election/export", handler.GetExportResult)

	// Web Socket Vote Stream
	ws := app.Group("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	ws.Get("/candidates", websocket.New(handler.CandidatesVoteStream))
	ws.Get("/candidates/:candidateId", websocket.New(handler.CandidateVoteStream))

	// Template
	app.Get("/chart", handler.GetTopTenChart)
}
