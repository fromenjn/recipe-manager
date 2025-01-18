package handlers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/fromenjn/recipe-manager/internal/usecase"
)

type RecipeHandler struct {
	getRecipeUC     usecase.GetRecipeUseCase
	getAllRecipesUC usecase.GetAllRecipesUseCase
}

func NewRecipeHandler(getRecipeUC usecase.GetRecipeUseCase, getAllRecipesUC usecase.GetAllRecipesUseCase) *RecipeHandler {
	return &RecipeHandler{
		getRecipeUC:     getRecipeUC,
		getAllRecipesUC: getAllRecipesUC,
	}
}

// GetRecipe godoc
// @Summary      Retrieve a single recipe
// @Description  Get a recipe by its ID. Optionally, scale ingredient quantities by specifying `ingredient` and `quantity`.
// @Tags         recipes
// @Param        recipeID    path      string  true  "Recipe ID (e.g. '123')"
// @Param        ingredient  query     string  false "Ingredient to scale (e.g. 'Flour')"
// @Param        quantity    query     number  false "Quantity to scale the ingredient to (e.g. '300')"
// @Produce      json
// @Success      200  {object}  domain.Recipe
// @Failure      400  {string}  string "invalid 'quantity' query parameter"
// @Failure      404  {string}  string "recipe not found"
// @Failure      500  {string}  string "internal server error"
// @Router       /recipes/{recipeID} [get]
func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	// e.g. /recipes/123?ingredient=Flour&quantity=300
	path := r.URL.Path
	// This is not robust - you'd use a proper router in practice
	recipeID := path[len("/recipes/"):]

	// Query params
	ingredient := r.URL.Query().Get("ingredient")
	quantityStr := r.URL.Query().Get("quantity")

	var quantity float64
	if quantityStr != "" {
		parsedQ, err := strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			http.Error(w, "invalid 'quantity' query parameter", http.StatusBadRequest)
			return
		}
		quantity = parsedQ
	}

	recipe, err := rh.getRecipeUC.Execute(recipeID, ingredient, quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Convert to JSON and return
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

// ListRecipes godoc
// @Summary      List all recipes
// @Description  Returns all recipes in the system
// @Tags         recipes
// @Produce      json
// @Success      200  {array}  domain.Recipe
// @Failure      500  {string}  string "failed to write response"
// @Router       /recipes [get]
func (rh *RecipeHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Listing all recipes")
	recipes, err := rh.getAllRecipesUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipes); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
