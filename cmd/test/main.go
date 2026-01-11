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
	repo, err := repository.NewSQLiteRecipesRepo("recipes.db")
	if err != nil {
		log.Fatal("Error al abrir la DB:", err)
	}

	// Crear el service
	svc := service.NewRecipeService(repo)

	// Probar ListRecipes
	recipes, err := svc.ListRecipes()
	if err != nil {
		log.Fatal("Error al listar recetas:", err)
	}

	// Mostrar resultados
	for _, r := range recipes {
		fmt.Printf("ID: %s, Name: %s, Desc: %s, Qty: %d %s, Difficulty: %s\n",
			r.ID, r.Name, r.Description, r.Quantity, r.Unit, r.Difficulty)
	}
}
