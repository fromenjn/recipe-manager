package usecase_test

import (
	"errors"
	"testing"

	"github.com/fromenjn/recipe-manager/internal/domain"
)

// Test ListAll success case
func TestMockRepo_ListAll_Success(t *testing.T) {
	// Prepare mock data
	mockData := map[string]domain.Recipe{
		"1": {ID: "1", Name: "Pancakes"},
		"2": {ID: "2", Name: "Omelette"},
	}

	// Create a mockRepo with recipes and no error
	repo := &mockRepo{
		recipes: mockData,
		err:     nil,
	}

	// Call ListAll
	recipes, err := repo.ListAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify result
	if len(recipes) != 2 {
		t.Fatalf("expected 2 recipes, got %d", len(recipes))
	}

	// We can convert the map to a slice to compare, or just check IDs/names
	//expected := []domain.Recipe{{ID: "1", Name: "Pancakes"}, {ID: "2", Name: "Omelette"}}

	// The order might vary if you want to check slices directly, so you could sort them or handle it individually.
	// For simplicity, let's do a reflect.DeepEqual check after sorting or just ignoring order:

	// A straightforward approach is to store them in a map for comparison, or just check individually:
	found := make(map[string]bool)
	for _, r := range recipes {
		found[r.ID] = true
	}
	if !found["1"] || !found["2"] {
		t.Errorf("missing expected recipes: got %v", recipes)
	}

	// If you want to do a direct slice comparison in order, you'd ensure a consistent ordering or just manually check.
}

// Test ListAll error case
func TestMockRepo_ListAll_Error(t *testing.T) {
	repo := &mockRepo{
		recipes: nil,
		err:     errors.New("some error"),
	}

	_, err := repo.ListAll()
	if err == nil {
		t.Error("expected error, got none")
	}
	if err.Error() != "some error" {
		t.Errorf("expected 'some error', got '%v'", err)
	}
}

// Test FindByID success case
func TestMockRepo_FindByID_Success(t *testing.T) {
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{
			"1": {ID: "1", Name: "Pancakes"},
		},
	}

	recipe, err := repo.FindByID("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if recipe == nil {
		t.Fatal("expected a recipe, got nil")
	}
	if recipe.ID != "1" || recipe.Name != "Pancakes" {
		t.Errorf("unexpected recipe data: %#v", recipe)
	}
}

// Test FindByID not found case
func TestMockRepo_FindByID_NotFound(t *testing.T) {
	repo := &mockRepo{
		recipes: map[string]domain.Recipe{
			"1": {ID: "1", Name: "Pancakes"},
		},
	}

	_, err := repo.FindByID("999")
	if err == nil {
		t.Error("expected an error for missing recipe, got none")
	}
}
