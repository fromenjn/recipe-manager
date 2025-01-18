package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/fromenjn/recipe-manager/internal/config"
	"github.com/fromenjn/recipe-manager/internal/domain"
	"github.com/fromenjn/recipe-manager/internal/handlers"
	"github.com/fromenjn/recipe-manager/internal/repository"
	"github.com/fromenjn/recipe-manager/internal/usecase"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	configPath := flag.String("config", "config.json", "Path to configuration JSON file")
	flag.Parse()
	// Load config
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize repository based on config
	repo, err := repository.NewJSONRepository(cfg.RecipesPath)
	if err != nil {
		log.Fatalf("Failed to init JSON repository: %v", err)
	}
	// Initialize domain service
	recipeService := domain.NewRecipeService()

	// Initialize use cases
	getRecipeUC := usecase.NewGetRecipeUseCase(repo, recipeService)
	getAllRecipesUC := usecase.NewGetAllRecipesUseCase(repo)

	// Initialize handlers
	recipeHandler := handlers.NewRecipeHandler(getRecipeUC, getAllRecipesUC)

	// Create router
	router := handlers.NewRouter(recipeHandler)

	// Start server on the configured port
	slog.Info(fmt.Sprintf("Starting server on %s", cfg.ServerPort))
	if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
		slog.Error(fmt.Sprintf("server error: %v", err))
	}
}
