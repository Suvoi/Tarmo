package service

import (
	"tarmo/internal/modules/recipes/model"
	"tarmo/internal/modules/recipes/port"
)

type RecipeService struct {
	repo port.RecipePort
}

func NewRecipeService(repo port.RecipePort) *RecipeService {
	return &RecipeService{repo: repo}
}

func (s *RecipeService) ListRecipes() ([]*model.Recipe, error) {
	return s.repo.GetAll()
}
