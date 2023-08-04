package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/meal-planner/api"
	"github.com/meal-planner/dbs"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if err := dbs.ConnectDb(); err != nil {
		panic(err)
	}

	api := api.NewRESTAPI()

	api.Serve(":8080")
}
