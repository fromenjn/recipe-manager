package usecase

import (
	"github.com/fromenjn/recipe-manager/internal/domain"
	"github.com/fromenjn/recipe-manager/internal/repository"
)

type GetRecipeUseCase interface {
	Execute(recipeID, ingredientConstraint string, quantityConstraint float64) (*domain.Recipe, error)
}

type getRecipeUseCase struct {
	repo    repository.RecipeRepository
	service domain.RecipeService
}

func NewGetRecipeUseCase(
	repo repository.RecipeRepository,
	service domain.RecipeService,
) GetRecipeUseCase {
	return &getRecipeUseCase{
		repo:    repo,
		service: service,
	}
}

// Execute fetches a recipe by ID and computes any ingredient ratios if requested.
func (uc *getRecipeUseCase) Execute(
	recipeID, ingredientConstraint string, quantityConstraint float64,
) (*domain.Recipe, error) {
	recipe, err := uc.repo.FindByID(recipeID)
	if err != nil {
		return nil, err
	}

	// Apply ratio logic if constraints are provided
	if err := uc.service.ComputeRatios(recipe, ingredientConstraint, quantityConstraint); err != nil {
		return nil, err
	}

	return recipe, nil
}
