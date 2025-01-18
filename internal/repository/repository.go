package repository

import "github.com/fromenjn/recipe-manager/internal/domain"

type RecipeRepository interface {
	FindByID(id string) (*domain.Recipe, error)
	ListAll() ([]domain.Recipe, error)
}
