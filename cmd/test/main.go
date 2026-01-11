package main

import (
	"fmt"
	"log"
	"tarmo/internal/adapters/repository/recipes"
	"tarmo/internal/core/recipes/ports"
	"tarmo/internal/core/recipes/services"
)

func main() {
	var repo ports.RecipePort
	repo, err := recipes.NewSQLiteRecipesRepo("data/test.db")
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	svc := services.NewRecipeService(repo)

	recipes, err := svc.ListRecipes()
	if err != nil {
		log.Fatal("Error al listar recetas:", err)
	}

	for _, r := range recipes {
		fmt.Printf("ID: %d, Name: %s, Desc: %s",
			r.ID, r.Name, r.Description)
	}
}
