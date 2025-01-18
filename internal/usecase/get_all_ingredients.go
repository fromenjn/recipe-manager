package usecase

import (
	"github.com/fromenjn/recipe-manager/internal/repository"
)

type GetAllIngredientsUseCase interface {
	Execute() ([]string, error)
}

type getAllIngredientsUseCase struct {
	repo repository.RecipeRepository
}

func NewGetAllIngredientsUseCase(repo repository.RecipeRepository) GetAllIngredientsUseCase {
	return &getAllIngredientsUseCase{
		repo: repo,
	}
}

// Execute returns all Ingredients from the repository.
func (uc *getAllIngredientsUseCase) Execute() ([]string, error) {
	recipes, err := uc.repo.ListAll()
	if err != nil {
		return nil, err
	}
	allIngredients := make([]string, 0)
	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			mustAdd := true
			for _, itr := range allIngredients {
				if ingredient.Name == itr {
					mustAdd = false
				}
			}
			if mustAdd {
				allIngredients = append(allIngredients, ingredient.Name)
			}
		}
	}
	return allIngredients, nil
}
