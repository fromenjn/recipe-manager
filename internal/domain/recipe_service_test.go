package domain

import (
	"testing"
)

func TestComputeRatios_NoConstraint(t *testing.T) {
	service := NewRecipeService()

	recipe := &Recipe{
		ID:   "test",
		Name: "Pancakes",
		Ingredients: []Ingredient{
			{Name: "Flour", Quantity: 200, Unit: "grams"},
			{Name: "Milk", Quantity: 300, Unit: "ml"},
		},
	}

	// No constraint name or invalid quantity => no scaling
	err := service.ComputeRatios(recipe, "", 0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check that quantities remain unchanged
	if recipe.Ingredients[0].Quantity != 200 {
		t.Errorf("expected flour quantity 200, got %v", recipe.Ingredients[0].Quantity)
	}
	if recipe.Ingredients[1].Quantity != 300 {
		t.Errorf("expected milk quantity 300, got %v", recipe.Ingredients[1].Quantity)
	}
}

func TestComputeRatios_WithConstraint(t *testing.T) {
	service := NewRecipeService()

	recipe := &Recipe{
		ID:   "test",
		Name: "Pancakes",
		Ingredients: []Ingredient{
			{Name: "Flour", Quantity: 200, Unit: "grams"},
			{Name: "Milk", Quantity: 300, Unit: "ml"},
		},
	}

	// Scale flour to 400 grams
	err := service.ComputeRatios(recipe, "Flour", 400)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// We expect ratio = 400 / 200 = 2
	// So Flour => 400, Milk => 600
	if recipe.Ingredients[0].Quantity != 400 {
		t.Errorf("expected flour quantity 400, got %v", recipe.Ingredients[0].Quantity)
	}
	if recipe.Ingredients[1].Quantity != 600 {
		t.Errorf("expected milk quantity 600, got %v", recipe.Ingredients[1].Quantity)
	}
}

func TestComputeRatios_IngredientNotFound(t *testing.T) {
	service := NewRecipeService()

	recipe := &Recipe{
		ID:   "test",
		Name: "Pancakes",
		Ingredients: []Ingredient{
			{Name: "Flour", Quantity: 200, Unit: "grams"},
			{Name: "Milk", Quantity: 300, Unit: "ml"},
		},
	}

	err := service.ComputeRatios(recipe, "Sugar", 100)
	if err == nil {
		t.Error("expected error due to missing ingredient, got none")
	}
}
