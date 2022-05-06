package database

import (
	"fmt"
	"line-town-election-api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Instance
var Database *gorm.DB

func SetupDatabase() {
	//Open Database
	var err error
	Database, err = gorm.Open(sqlite.Open("./database/election.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to conenct database")
	}

	//Auto Migration
	err = Database.AutoMigrate(
		&model.Candidate{},
		&model.Vote{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	//Success
	fmt.Println("Database Migrated")
}
