package main

import (
	"github.com/meal-planner/api"
	"github.com/meal-planner/dbs"
)

func main() {
	if err := dbs.ConnectDb(); err != nil {
		panic(err)
	}

	api := api.NewRESTAPI()
	api.Serve(":8080")
}
