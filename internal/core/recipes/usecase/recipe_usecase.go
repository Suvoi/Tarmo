package usecase

import (
	"fmt"
	"tarmo/internal/core/recipes/domain"
	ports "tarmo/internal/core/recipes/ports/outbound"
)

type RecipeUseCase struct {
	port ports.RecipePort
}

func NewRecipeUseCase(port ports.RecipePort) *RecipeUseCase {
	return &RecipeUseCase{port: port}
}

func (s *RecipeUseCase) ListRecipes() ([]*domain.Recipe, error) {
	return s.port.GetAll()
}

func (s *RecipeUseCase) GetRecipe(id int) (*domain.Recipe, error) {
	return s.port.GetByID(id)
}

func (s *RecipeUseCase) CreateRecipe(rcp *domain.Recipe) error {

	// Basic Validations
	if rcp.Name == "" {
		return fmt.Errorf("name is required")
	}

	if rcp.Quantity == 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	if rcp.Unit == "" {
		return fmt.Errorf("unit must be defined")
	}

	if rcp.Difficulty <= 0 || rcp.Difficulty > 5 {
		return fmt.Errorf("difficulty must be between 0 and 5")
	}

	return s.port.Create(rcp)
}

func (s *RecipeUseCase) DeleteRecipe(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid recipe id")
	}

	return s.port.Delete(id)
}
