package repository

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewJSONRepository_SingleFile(t *testing.T) {
	// Create a temporary directory
	dir, err := os.MkdirTemp("", "test-recipes-")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir) // clean up

	// Write a test JSON file
	testFilePath := filepath.Join(dir, "recipes1.json")
	content := []byte(`
        {
            "id": "1",
            "name": "Pancakes",
            "ingredients": [
                {"name": "Flour", "quantity": 200, "unit": "grams"},
                {"name": "Milk", "quantity": 300, "unit": "ml"}
            ]
        }
    `)
	if err := os.WriteFile(testFilePath, content, 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	repo, err := NewJSONRepository(dir)
	if err != nil {
		t.Fatalf("failed to create repository: %v", err)
	}

	// Check if we can retrieve the recipe
	rcp, err := repo.FindByID("1")
	if err != nil {
		t.Errorf("expected to find recipe '1', got error %v", err)
	}
	if rcp == nil {
		t.Fatal("expected a recipe, got nil")
	}
	if rcp.Name != "Pancakes" {
		t.Errorf("expected 'Pancakes', got %s", rcp.Name)
	}
	if len(rcp.Ingredients) != 2 {
		t.Errorf("expected 2 ingredients, got %d", len(rcp.Ingredients))
	}
}

func TestNewJSONRepository_MultipleFiles(t *testing.T) {
	// Create a temporary directory
	dir, err := os.MkdirTemp("", "test-recipes-")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	// Write multiple JSON files
	file1 := filepath.Join(dir, "recipes1.json")
	file2 := filepath.Join(dir, "recipes2.json")

	content1 := []byte(`
        {
            "id": "1",
            "name": "Pancakes",
            "ingredients": [
                {"name": "Flour", "quantity": 200, "unit": "grams"}
            ]
        }
    `)

	content2 := []byte(`
        {
            "id": "2",
            "name": "Omelette",
            "ingredients": [
                {"name": "Egg", "quantity": 3, "unit": "pieces"}
            ]
        }
    `)

	if err := os.WriteFile(file1, content1, 0644); err != nil {
		t.Fatalf("failed to write file1: %v", err)
	}
	if err := os.WriteFile(file2, content2, 0644); err != nil {
		t.Fatalf("failed to write file2: %v", err)
	}

	repo, err := NewJSONRepository(dir)
	if err != nil {
		t.Fatalf("failed to create repository from directory: %v", err)
	}

	// Check first recipe
	rcp1, err := repo.FindByID("1")
	if err != nil {
		t.Errorf("expected to find recipe '1', got error %v", err)
	}
	if rcp1 != nil && rcp1.Name != "Pancakes" {
		t.Errorf("expected 'Pancakes', got %s", rcp1.Name)
	}

	// Check second recipe
	rcp2, err := repo.FindByID("2")
	if err != nil {
		t.Errorf("expected to find recipe '2', got error %v", err)
	}
	if rcp2 != nil && rcp2.Name != "Omelette" {
		t.Errorf("expected 'Omelette', got %s", rcp2.Name)
	}
}
