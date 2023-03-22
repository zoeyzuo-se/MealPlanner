package models

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	RecipeName            string `json:"recipe_name"`
	RecipePreparationTime int    `json:"recipe_preparation_time"`
	RecipeCategory        string `json:"recipe_category"`
	RecipeMethod          string `json:"recipe_method"`
}
