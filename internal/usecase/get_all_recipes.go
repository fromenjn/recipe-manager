package usecase

import (
	"github.com/fromenjn/recipe-manager/internal/domain"
	"github.com/fromenjn/recipe-manager/internal/repository"
)

type GetAllRecipesUseCase interface {
	Execute() ([]domain.Recipe, error)
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
func (uc *getAllRecipesUseCase) Execute() ([]domain.Recipe, error) {
	return uc.repo.ListAll()
}
