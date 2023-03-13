package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meal-planner/models"
)

var recipes []models.Recipe

func addRecipe(c *gin.Context) {
	var newRecipe models.Recipe

	// Bind the JSON request body to the Recipe struct
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new RecipeID by adding 1 to the length of the meals slice
	newRecipe.ID = len(recipes) + 1

	// Append the new Recipe to the meals slice
	recipes = append(recipes, newRecipe)

	// Marshal the meals slice to JSON and write it to a file
	data, _ := json.Marshal(recipes)
	err := ioutil.WriteFile("recipes.json", data, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Return the inserted Recipe
	c.JSON(http.StatusOK, gin.H{"data": newRecipe})
}

func getRecipeByName(c *gin.Context) {
	name := c.Param("name")
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	var payload []models.Recipe
	err = json.Unmarshal(content, &payload)
	for _, a := range payload {
		if a.RecipeName == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
