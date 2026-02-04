package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"tarmo/internal/adapters/inbound/rest/middleware"
	"tarmo/internal/adapters/inbound/rest/recipes"
	sqliteRecipes "tarmo/internal/adapters/outbound/persistence/sqlite"
	"tarmo/internal/core/recipes/usecase"
	"tarmo/internal/lib/logger"
)

func main() {
	logger.Info("Starting Tarmo rest API on port 9136")

	// DB adapter
	repo, err := sqliteRecipes.NewSQLiteRecipesRepo("data/db.sqlite")
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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/", func(r chi.Router) {
		recipes.RegisterRoutes(r, handler)
	})

	err = http.ListenAndServe(":9136", r)
	if err != nil {
		logger.Fatal("server error: %v", err)
		return
	}

}
