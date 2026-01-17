package usecase

import (
	"fmt"
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/ports/outbound"
)

type RecipeUseCase struct {
	port outbound.RecipeRepositoryPort
}

func NewRecipeUseCase(port outbound.RecipeRepositoryPort) *RecipeUseCase {
	return &RecipeUseCase{port: port}
}

func (uc *RecipeUseCase) ListRecipes() ([]*domain.Recipe, error) {
	return uc.port.FindAll()
}

func (uc *RecipeUseCase) GetRecipe(id int) (*domain.Recipe, error) {
	return uc.port.FindByID(id)
}

func (uc *RecipeUseCase) CreateRecipe(rcp *domain.Recipe) error {

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

	return uc.port.Save(rcp)
}

func (uc *RecipeUseCase) DeleteRecipe(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid recipe id")
	}

	return uc.port.Remove(id)
}
