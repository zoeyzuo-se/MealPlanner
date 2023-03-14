package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/meal-planner/models"
	"github.com/meal-planner/pkg/recipes"
)

var recipe_list []models.Recipe

func AddRecipe(c *gin.Context) {
	var newRecipe models.Recipe

	// Bind the JSON request body to the Recipe struct
	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new RecipeID by adding 1 to the length of the recipes slice
	newRecipe.ID = len(recipe_list) + 1
	categoryError := fmt.Sprintf("Recipe catogory should be one of: %q", strings.Join(recipes.GetRecipeCategories(), " "))

	if !recipes.IsRecipeCategoryValid(newRecipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": categoryError})
		return
	}

	methodError := fmt.Sprintf("Recipe method should be one of: %q", strings.Join(recipes.GetRecipeMethods(), " "))
	if !recipes.IsRecipeMethodValid(newRecipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": methodError})
		return
	}

	// Append the new Recipe to the recipes slice
	recipe_list = append(recipe_list, newRecipe)

	// Marshal the recipes slice to JSON and write it to a file
	data, _ := json.Marshal(recipe_list)
	err := ioutil.WriteFile("recipes.json", data, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Return the inserted Recipe
	c.JSON(http.StatusOK, gin.H{"data": newRecipe})
}

func GetRecipeByName(c *gin.Context) {
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
