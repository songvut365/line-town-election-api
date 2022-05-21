package main

import (
	"line-town-election-api/database"
	"line-town-election-api/handler"
	"line-town-election-api/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

// Global Variable
var ElectionStatus = true

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./template", ".html"),
	})

	//Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	//Setup
	SetupEnv()
	router.SetupRouter(app)
	database.SetupDatabase()

	//Inject global variable
	handler.ElectionStatus = ElectionStatus

	//Run
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
