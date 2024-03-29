package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meal-planner/models"
	"github.com/meal-planner/pkg/recipes"
)

type RESTApiv1 struct {
	router *gin.Engine
}

func NewRESTAPI() *RESTApiv1 {
	router := gin.Default()
	api := &RESTApiv1{router}

	router.POST("/recipe", api.AddRecipe)
	router.GET("/recipe/:name", api.GetRecipeByName)
	router.GET("/recipes", api.GetRecipes)
	router.PUT("/recipe/:name", api.UpdateRecipe)
	router.DELETE("/recipe/:name", api.DeleteRecipe)

	return api
}

func (api *RESTApiv1) Serve(addr string) error {
	return api.router.Run(addr)
}

func (api *RESTApiv1) AddRecipe(c *gin.Context) {
	var newRecipe models.Recipe

	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := recipes.GetRecipeByName(newRecipe.RecipeName, &newRecipe); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe already exists"})
		return
	}

	if !recipes.IsRecipeCategoryValid(newRecipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": recipes.GenerateCategoryError()})
		return
	}

	if !recipes.IsRecipeMethodValid(newRecipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": recipes.GenerateMethodError()})
		return
	}

	if err := recipes.AddRecipe(&newRecipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add recipe"})
	}

	c.JSON(http.StatusOK, gin.H{"data": newRecipe})
}

func (api *RESTApiv1) GetRecipeByName(c *gin.Context) {
	name := c.Param("name")
	var recipe models.Recipe

	if err := recipes.GetRecipeByName(name, &recipe); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func (api *RESTApiv1) GetRecipes(c *gin.Context) {
	var recipeList []models.Recipe

	if err := recipes.GetRecipes(&recipeList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recipes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipeList})
}

func (api *RESTApiv1) UpdateRecipe(c *gin.Context) {
	name := c.Param("name")
	var recipe models.Recipe

	if err := recipes.GetRecipeByName(name, &recipe); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !recipes.IsRecipeCategoryValid(recipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": recipes.GenerateCategoryError()})
		return
	}

	if !recipes.IsRecipeMethodValid(recipe) {
		c.JSON(http.StatusBadRequest, gin.H{"error": recipes.GenerateMethodError()})
		return
	}

	if err := recipes.UpdateRecipe(name, recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func (api *RESTApiv1) DeleteRecipe(c *gin.Context) {
	name := c.Param("name")
	var recipe models.Recipe

	if err := recipes.GetRecipeByName(name, &recipe); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	if err := recipes.DeleteRecipe(&recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": name + " successfully deleted."})
}

func (api *RESTApiv1) GetRecipeByCategory(c *gin.Context) {
	category := c.Param("category")
	var recipeList []models.Recipe

	if err := recipes.GetRecipeByCategory(category, &recipeList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recipes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipeList})
}
