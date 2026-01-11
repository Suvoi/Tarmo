package main

import (
	"fmt"
	"log"
	"tarmo/internal/modules/recipes/port"
	"tarmo/internal/modules/recipes/service"
	"tarmo/internal/repository"
)

func main() {
	var repo port.RecipePort
	repo, err := repository.NewSQLiteRecipesRepo("data/test.db")
	if err != nil {
		log.Fatal("Error al abrir la DB:", err)
	}

	svc := service.NewRecipeService(repo)

	recipes, err := svc.ListRecipes()
	if err != nil {
		log.Fatal("Error al listar recetas:", err)
	}

	for _, r := range recipes {
		fmt.Printf("ID: %s, Name: %s, Desc: %s, Qty: %d %s, Difficulty: %s\n",
			r.ID, r.Name, r.Description, r.Quantity, r.Unit, r.Difficulty)
	}
}
