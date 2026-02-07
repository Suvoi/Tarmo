package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"tarmo/internal/adapters/inbound/rest/middleware"
	"tarmo/internal/adapters/inbound/rest/recipes"
	sqliteRecipes "tarmo/internal/adapters/outbound/persistence/sqlite"
	"tarmo/internal/config"
	"tarmo/internal/core/recipes/usecase"
	"tarmo/internal/lib/logger"
)

// @title Tarmo API
// @version 2.1.0
// @description Optimize and control batches based on recipes.
// @host localhost:9136
// @BasePath /
func main() {
	cfg := config.Load()

	logger.Info("Starting Tarmo on port %s", cfg.Port)

	// DB adapter
	repo, err := sqliteRecipes.NewSQLiteRecipesRepo(cfg.DBPath)
	if err != nil {
		logger.Fatal("failed to connect db: %v", err)
		return
	}

	// Usecase
	uc := usecase.NewRecipeUseCase(repo)

	// Handler REST
	handler := recipes.NewHandler(uc)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logging)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/", func(r chi.Router) {
		recipes.RegisterRoutes(r, handler)
	})

	err = http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		logger.Fatal("server error: %v", err)
		return
	}

}
