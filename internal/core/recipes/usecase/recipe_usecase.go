package usecase

import (
	"errors"
	"fmt"
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/ports/outbound"
)

type CreateRecipeCommand struct {
	Name        string
	Description string
	Quantity    int
	Unit        string
	Difficulty  int
	Steps       []CreateStepCommand
}

type CreateStepCommand struct {
	Name         string
	Instructions string
}

type RecipeUseCase struct {
	port outbound.RecipeRepositoryPort
}

func NewRecipeUseCase(port outbound.RecipeRepositoryPort) *RecipeUseCase {
	return &RecipeUseCase{port: port}
}

func (uc *RecipeUseCase) GetAll() ([]*domain.Recipe, error) {
	return uc.port.FindAll()
}

func (uc *RecipeUseCase) GetByID(id int) (*domain.Recipe, error) {
	return uc.port.FindByID(id)
}

func (uc *RecipeUseCase) Create(cmd CreateRecipeCommand) error {

	// Basic Validations
	if cmd.Name == "" {
		return errors.New("name is required")
	}

	if cmd.Quantity == 0 {
		return errors.New("quantity must be greater than 0")
	}

	if cmd.Unit == "" {
		return errors.New("unit must be defined")
	}

	if cmd.Difficulty < 0 || cmd.Difficulty > 5 {
		return errors.New("difficulty must be between 0 and 5")
	}

	if len(cmd.Steps) == 0 {
		return errors.New("recipe must have at least one step")
	}

	steps := make([]domain.Step, 0, len(cmd.Steps))
	for i, s := range cmd.Steps {
		if s.Name == "" {
			return fmt.Errorf("step %d name is required", i+1)
		}

		steps = append(steps, domain.Step{
			Order:        i + 1,
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}

	recipe := &domain.Recipe{
		Name:        cmd.Name,
		Description: cmd.Description,
		Quantity:    cmd.Quantity,
		Unit:        cmd.Unit,
		Difficulty:  cmd.Difficulty,
		Steps:       steps,
	}

	return uc.port.Save(recipe)
}

func (uc *RecipeUseCase) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid recipe id")
	}

	return uc.port.Remove(id)
}
