package database

import (
	"line-town-election-api/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Instance
var Database *gorm.DB

func SetupDatabase() {
	// Switch MySQL and SQLite3
	databaseType := os.Getenv("DATABASE_TYPE")
	var err error

	//Open Database
	switch databaseType {
	case "MySQL":
		Database, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_URI")), &gorm.Config{})
	case "SQLite":
		Database, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_URI")), &gorm.Config{})
	}

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
