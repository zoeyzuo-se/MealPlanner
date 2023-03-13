package recipes

import "github.com/meal-planner/models"

var recipeCategories = []string{"荤", "素"}
var recipeMethods = []string{"炒", "蒸", "炖", "烤"}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsRecipeCategoryValid(input models.Recipe) bool {
	return stringInSlice(input.RecipeCategory, recipeCategories)
}

func IsRecipeMethodValid(input models.Recipe) bool {
	return stringInSlice(input.RecipeMethod, recipeMethods)
}

func GetRecipeCategories() []string {
	return recipeCategories
}

func GetRecipeMethods() []string {
	return recipeMethods
}
