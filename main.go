package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meal-planner/api"
)

func main() {
	router := gin.Default()

	router.POST("/recipe", api.AddRecipe)
	router.GET("/recipe/:name", api.GetRecipeByName)

	router.Run(":8080")
}
