package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fromenjn/recipe-manager/internal/domain"
)

type jsonRepository struct {
	dirPath string
	recipes map[string]domain.Recipe
}

// NewJSONRepository creates a new repository that reads from all JSON files in a directory.
func NewJSONRepository(dirPath string) (RecipeRepository, error) {
	repo := &jsonRepository{
		dirPath: dirPath,
		recipes: make(map[string]domain.Recipe),
	}

	if err := repo.loadRecipes(); err != nil {
		return nil, err
	}

	return repo, nil
}

// loadRecipes reads all .json files in dirPath, accumulates them in a map by ID.
func (r *jsonRepository) loadRecipes() error {
	// Verify the directory exists
	info, err := os.Stat(r.dirPath)
	if err != nil {
		return fmt.Errorf("failed to open directory %s: %w", r.dirPath, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path %s is not a directory", r.dirPath)
	}

	// Walk the directory for .json files
	err = filepath.Walk(r.dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If it's a directory, skip it (except the root)
		if info.IsDir() && path != r.dirPath {
			return filepath.SkipDir
		}

		// Only parse .json files
		if filepath.Ext(path) == ".json" {
			if loadErr := r.parseFile(path); loadErr != nil {
				// Return an error so we fail fast on a broken file
				return loadErr
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed walking directory: %w", err)
	}

	return nil
}

// parseFile reads a single JSON file and accumulates recipe data into the repository map.
func (r *jsonRepository) parseFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	var fileRecipes []domain.Recipe
	if err := json.Unmarshal(data, &fileRecipes); err != nil {
		return fmt.Errorf("failed to unmarshal JSON in file %s: %w", path, err)
	}

	for _, rcp := range fileRecipes {
		if _, exists := r.recipes[rcp.ID]; exists {
			// Warn or handle duplicates as you see fit
			// For example, you could override the existing recipe, or skip
			return fmt.Errorf("duplicate recipe ID '%s' found in file %s", rcp.ID, path)
		}
		r.recipes[rcp.ID] = rcp
	}

	return nil
}

func (r *jsonRepository) FindByID(id string) (*domain.Recipe, error) {
	recipe, ok := r.recipes[id]
	if !ok {
		return nil, errors.New("recipe not found")
	}
	// Return a copy to avoid side-effects on the in-memory map
	return &recipe, nil
}

// ListAll returns all recipes in the repository.
func (r *jsonRepository) ListAll() ([]domain.Recipe, error) {
	recipes := make([]domain.Recipe, 0, len(r.recipes))
	for _, rcp := range r.recipes {
		recipes = append(recipes, rcp)
	}
	return recipes, nil
}
