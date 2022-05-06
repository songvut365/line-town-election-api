package database

import (
	"fmt"
	"line-town-election-api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Instance
var DB *gorm.DB

func SetupDatabase() {
	//Open Database
	db, err := gorm.Open(sqlite.Open("./database/election.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to conenct database")
	}

	//Auto Migration
	err = db.AutoMigrate(
		&model.Candidate{},
		&model.Vote{},
		&model.System{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	//Success
	fmt.Println("Database Migrated")
}
