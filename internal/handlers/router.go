package handlers

import (
	"net/http"
)

func NewRouter(recipeHandler *RecipeHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/recipe/", recipeHandler.GetRecipe)
	mux.HandleFunc("/recipes", recipeHandler.ListRecipes)
	mux.HandleFunc("/ingredients", recipeHandler.ListIngredients)

	return mux
}
