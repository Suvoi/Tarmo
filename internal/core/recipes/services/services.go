package services

import (
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/ports"
)

type RecipeService struct {
	port ports.RecipePort
}

func NewRecipeService(port ports.RecipePort) *RecipeService {
	return &RecipeService{port: port}
}

func (s *RecipeService) ListRecipes() ([]*domain.Recipe, error) {
	return s.port.GetAll()
}
