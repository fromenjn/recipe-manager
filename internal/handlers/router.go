package handlers

import (
	"net/http"
)

func NewRouter(recipeHandler *RecipeHandler) http.Handler {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./dist"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/recipe/", recipeHandler.GetRecipe)
	mux.HandleFunc("/recipes", recipeHandler.ListRecipes)
	mux.HandleFunc("/ingredients", recipeHandler.ListIngredients)

	muxWithCors := WithCORS(mux)
	return muxWithCors
}
