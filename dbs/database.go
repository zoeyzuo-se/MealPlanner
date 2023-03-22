package dbs

import (
	"log"

	"github.com/meal-planner/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database \n", err.Error())
		return err
	}

	log.Println("Database connection successfully opened")
	db.AutoMigrate(&models.Recipe{})
	Database.Db = db
	return nil
}
