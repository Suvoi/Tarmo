package main

import (
	"fmt"
	"log"
	"tarmo/internal/adapters/outbound/repository/recipes"
	"tarmo/internal/core/recipes/ports/outbound"
	"tarmo/internal/core/recipes/usecase"
)

func main() {
	var repo outbound.RecipePort
	repo, err := recipes.NewSQLiteRecipesRepo("data/test.db")
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	uc := usecase.NewRecipeUseCase(repo)

	recipes, err := uc.ListRecipes()
	if err != nil {
		log.Fatal("Error al listar recetas:", err)
	}

	for _, r := range recipes {
		fmt.Printf("ID: %d, Name: %s, Desc: %s",
			r.ID, r.Name, r.Description)
	}
}
