package recipes

import (
	"fmt"
	"strings"

	"github.com/meal-planner/dbs"
	"github.com/meal-planner/models"
)

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

func GenerateCategoryError() string {
	return fmt.Sprintf("Recipe catogory should be one of: %q", strings.Join(recipeCategories, " "))
}

func GenerateMethodError() string {
	return fmt.Sprintf("Recipe method should be one of: %q", strings.Join(recipeMethods, " "))
}

func AddRecipe(newRecipe *models.Recipe) error {
	result := dbs.Database.Db.Create(newRecipe)
	return result.Error
}

func GetRecipeByName(name string, recipe *models.Recipe) error {
	result := dbs.Database.Db.Where("recipe_name = ?", name).First(recipe)
	return result.Error
}

func GetRecipeByCategory(category string, recipes *[]models.Recipe) error {
	result := dbs.Database.Db.Where("recipe_category = ?", category).Find(recipes)
	return result.Error
}

func GetRecipes(recipes *[]models.Recipe) error {
	result := dbs.Database.Db.Find(recipes)
	return result.Error
}

func UpdateRecipe(name string, recipe models.Recipe) error {
	result := dbs.Database.Db.Model(&models.Recipe{}).Where("recipe_name = ?", name).Updates(recipe)
	return result.Error
}

func DeleteRecipe(recipe *models.Recipe) error {
	result := dbs.Database.Db.Delete(recipe)
	return result.Error
}
