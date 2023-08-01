package dbs

import (
	"log"
	"os"

	"github.com/meal-planner/models"
	"gorm.io/gorm"

	"fmt"

	"gorm.io/driver/sqlserver"

	_ "github.com/microsoft/go-mssqldb"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() error {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		dbHost, dbUsername, dbPassword, dbPort, dbName)

	// Create connection pool
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database \n", err.Error())
		return err
	}
	db.AutoMigrate(&models.Recipe{})
	Database.Db = db
	fmt.Printf("Connected!")
	return nil
}
