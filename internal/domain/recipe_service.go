package domain

import "errors"

// RecipeService defines domain-level operations.
type RecipeService interface {
	ComputeRatios(recipe *Recipe, constraintName string, constraintQuantity float64) error
}

// recipeService is a concrete implementation of RecipeService.
type recipeService struct{}

// NewRecipeService returns a new RecipeService.
func NewRecipeService() RecipeService {
	return &recipeService{}
}

// ComputeRatios scales ingredient quantities in a recipe based on a constraint (if provided).
func (s *recipeService) ComputeRatios(recipe *Recipe, constraintName string, constraintQuantity float64) error {
	if constraintName == "" || constraintQuantity <= 0 {
		// No ratio scaling required or invalid constraint.
		return nil
	}

	// Find the ingredient that matches the constraintName.
	var baseQuantity float64
	for _, ingredient := range recipe.Ingredients {
		if ingredient.Name == constraintName {
			baseQuantity = ingredient.Quantity
			break
		}
	}

	if baseQuantity == 0 {
		return errors.New("ingredient constraint not found in recipe")
	}

	ratio := constraintQuantity / baseQuantity
	if ratio <= 0 {
		return errors.New("invalid scaling ratio")
	}

	// Scale all ingredient quantities.
	for i := range recipe.Ingredients {
		recipe.Ingredients[i].Quantity *= ratio
	}

	return nil
}
