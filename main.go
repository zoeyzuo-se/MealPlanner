package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Meal struct {
	ID                  int    `json:"id"`
	MealName            string `json:"meal_name"`
	MealPreparationTime int    `json:"meal_preparation_time"`
	MealCategory        string `json:"meal_category"`
	MealMethod          string `json:"meal_method"`
}

var mealCategories = []string{"荤", "素"}
var mealMethods = []string{"炒", "蒸", "炖", "烤"}

var meals []Meal

func main() {
	router := gin.Default()

	router.POST("/recipe", api.addRecipe)
	router.GET("/recipe/:name", api.getRecipeByName)

	router.Run(":8080")
}

func addMeal(c *gin.Context) {
	var meal Meal

	// Bind the JSON request body to the Meal struct
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new MealID by adding 1 to the length of the meals slice
	meal.ID = len(meals) + 1
	categoryError := fmt.Sprintf("Meal catogory should be one of: %q", strings.Join(mealCategories, " "))

	if !isMealCategoryValid(meal) {
		c.JSON(http.StatusBadRequest, gin.H{"error": categoryError})
		return
	}

	methodError := fmt.Sprintf("Meal method should be one of: %q", strings.Join(mealMethods, " "))
	if !isMealMethodValid(meal) {
		c.JSON(http.StatusBadRequest, gin.H{"error": methodError})
		return
	}

	// Append the new Meal to the meals slice
	meals = append(meals, meal)

	// Marshal the meals slice to JSON and write it to a file
	data, _ := json.Marshal(meals)
	err := ioutil.WriteFile("meals.json", data, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Return the inserted Meal
	c.JSON(http.StatusOK, gin.H{"data": meal})
}

func getMealByName(c *gin.Context) {
	name := c.Param("name")
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile("meals.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	var payload []Meal
	err = json.Unmarshal(content, &payload)
	for _, a := range payload {
		if a.MealName == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isMealCategoryValid(input Meal) bool {
	return stringInSlice(input.MealCategory, mealCategories)
}

func isMealMethodValid(input Meal) bool {
	return stringInSlice(input.MealMethod, mealMethods)
}
