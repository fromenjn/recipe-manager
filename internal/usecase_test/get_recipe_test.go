package usecase_test

import (
	"errors"
	"testing"

	"github.com/fromenjn/recipe-manager/internal/domain"
	"github.com/fromenjn/recipe-manager/internal/usecase"
)

// mockRepo is a mock implementation of the repository.RecipeRepository.
type mockRepo struct {
	recipes map[string]domain.Recipe
	err     error
}

func (m *mockRepo) FindByID(id string) (*domain.Recipe, error) {
	if m.err != nil {
		return nil, m.err
	}
	r, ok := m.recipes[id]
	if !ok {
		return nil, errors.New("recipe not found")
	}
	return &r, nil
}

// ListAll returns all recipes in the map as a slice, or an error if err != nil.
func (m *mockRepo) ListAll() ([]domain.Recipe, error) {
	if m.err != nil {
		return nil, m.err
	}

	results := make([]domain.Recipe, 0, len(m.recipes))
	for _, recipe := range m.recipes {
		rCopy := recipe
		results = append(results, rCopy)
	}
	return results, nil
}

// mockService is a mock implementation of the domain.RecipeService.
type mockService struct {
	computeErr error
	lastCall   struct {
		recipe           *domain.Recipe
		constraintName   string
		constraintAmount float64
	}
}

func (m *mockService) ComputeRatios(recipe *domain.Recipe, constraintName string, constraintAmount float64) error {
	m.lastCall.recipe = recipe
	m.lastCall.constraintName = constraintName
	m.lastCall.constraintAmount = constraintAmount

	return m.computeErr
}

func TestGetRecipeUseCase_Execute_NoScaling(t *testing.T) {
	// Setup
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{
			"1": {
				ID:   "1",
				Name: "Pancakes",
			},
		},
	}
	service := &mockService{}

	uc := usecase.NewGetRecipeUseCase(repo, service)

	// Execute use case without constraints
	result, err := uc.Execute("1", "", 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify recipe is returned
	if result == nil {
		t.Fatal("expected a recipe, got nil")
	}
	if result.ID != "1" {
		t.Errorf("expected recipe ID '1', got '%s'", result.ID)
	}

	// Verify ComputeRatios was called with no constraints
	if service.lastCall.recipe.ID != "1" {
		t.Errorf("expected recipe ID to be '1', got '%s'", service.lastCall.recipe.ID)
	}
	if service.lastCall.constraintName != "" || service.lastCall.constraintAmount != 0 {
		t.Error("expected empty constraint name and 0 constraint amount")
	}
}

func TestGetRecipeUseCase_Execute_WithScaling(t *testing.T) {
	// Setup
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{
			"1": {
				ID:   "1",
				Name: "Pancakes",
			},
		},
	}
	service := &mockService{}

	uc := usecase.NewGetRecipeUseCase(repo, service)

	// Execute with constraints
	result, err := uc.Execute("1", "Flour", 500)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected a recipe, got nil")
	}

	// Verify constraints were passed to the service
	if service.lastCall.constraintName != "Flour" {
		t.Errorf("expected constraintName 'Flour', got '%s'", service.lastCall.constraintName)
	}
	if service.lastCall.constraintAmount != 500 {
		t.Errorf("expected constraintAmount 500, got %f", service.lastCall.constraintAmount)
	}
}

func TestGetRecipeUseCase_Execute_RecipeNotFound(t *testing.T) {
	// Setup
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{},
	}
	service := &mockService{}
	uc := usecase.NewGetRecipeUseCase(repo, service)

	// Execute with non-existent ID
	_, err := uc.Execute("999", "", 0)
	if err == nil {
		t.Error("expected error for missing recipe, got none")
	}
}

func TestGetRecipeUseCase_Execute_ServiceError(t *testing.T) {
	// Setup
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{
			"1": {
				ID:   "1",
				Name: "Pancakes",
			},
		},
	}
	service := &mockService{
		computeErr: errors.New("some ratio error"),
	}
	uc := usecase.NewGetRecipeUseCase(repo, service)

	_, err := uc.Execute("1", "Flour", 500)
	if err == nil {
		t.Error("expected error from service, got none")
	}
	if err.Error() != "some ratio error" {
		t.Errorf("expected 'some ratio error', got '%s'", err.Error())
	}
}
