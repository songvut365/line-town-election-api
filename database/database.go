package database

import (
	"line-town-election-api/model"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Instance
var Database *gorm.DB

func SetupDatabase() {
	//Open Database
	var err error
	Database, err = gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("Failed to conenct database")
	}

	//Auto Migration
	err = Database.AutoMigrate(
		&model.Candidate{},
		&model.Vote{},
		&model.LogVote{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	//Success
	log.Println("Database Migrated")
}
