package main

import (
	"line-town-election-api/database"
	"line-town-election-api/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	//Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	//Setup
	router.SetupRouter(*app)
	database.SetupDatabase()

	//Run
	log.Fatal(app.Listen(":8080"))
}
