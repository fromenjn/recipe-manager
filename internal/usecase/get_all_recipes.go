package usecase

import (
	"github.com/fromenjn/recipe-manager/internal/domain"
	"github.com/fromenjn/recipe-manager/internal/repository"
)

type GetAllRecipesUseCase interface {
	Execute(ingredientConstraint string) ([]domain.Recipe, error)
}

type getAllRecipesUseCase struct {
	repo repository.RecipeRepository
}

func NewGetAllRecipesUseCase(repo repository.RecipeRepository) GetAllRecipesUseCase {
	return &getAllRecipesUseCase{
		repo: repo,
	}
}

// Execute returns all recipes from the repository.
func (uc *getAllRecipesUseCase) Execute(ingredientConstraint string) ([]domain.Recipe, error) {
	recipes, err := uc.repo.ListAll()
	if err != nil {
		return nil, err
	}
	if ingredientConstraint != "" {
		newRecipes := make([]domain.Recipe, 0)
		for _, recipe := range recipes {
			for _, ingredient := range recipe.Ingredients {
				if ingredient.Name == ingredientConstraint {
					newRecipes = append(newRecipes, recipe)
				}
			}
		}
		return newRecipes, nil
	}
	return uc.repo.ListAll()
}
