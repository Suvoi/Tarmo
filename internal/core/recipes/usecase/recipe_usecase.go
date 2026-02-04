package usecase

import (
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/ports/inbound"
	"tarmo/internal/core/recipes/ports/outbound"
)

type RecipeUseCase struct {
	repo outbound.RecipeRepositoryPort
}

func NewRecipeUseCase(repo outbound.RecipeRepositoryPort) *RecipeUseCase {
	return &RecipeUseCase{repo: repo}
}

func (uc *RecipeUseCase) GetAll() ([]inbound.RecipeDTO, error) {
	recipes, err := uc.repo.FindAll()
	if err != nil {
		return []inbound.RecipeDTO{}, err
	}

	dtos := make([]inbound.RecipeDTO, 0, len(recipes))
	for _, recipe := range recipes {
		if recipe == nil {
			continue
		}

		dtos = append(dtos, inbound.RecipeDTO{
			ID:          recipe.ID(),
			Name:        recipe.Name(),
			Description: recipe.Description(),
			Quantity:    recipe.Quantity(),
			Unit:        recipe.Unit(),
			Difficulty:  recipe.Difficulty(),
		})
	}

	return dtos, nil
}

func (uc *RecipeUseCase) GetByID(id int) (inbound.RecipeDTO, error) {
	recipe, err := uc.repo.FindByID(id)
	if err != nil {
		return inbound.RecipeDTO{}, err
	}

	steps := make([]inbound.StepDTO, 0, len(recipe.Steps()))
	for _, step := range recipe.Steps() {
		steps = append(steps, inbound.StepDTO{
			Name:         step.Name(),
			Instructions: step.Instructions(),
			Order:        step.Order(),
		})
	}

	return inbound.RecipeDTO{
		ID:          recipe.ID(),
		Name:        recipe.Name(),
		Description: recipe.Description(),
		Quantity:    recipe.Quantity(),
		Unit:        recipe.Unit(),
		Difficulty:  recipe.Difficulty(),
		Steps:       steps,
	}, nil
}

func (uc *RecipeUseCase) Create(cmd inbound.CreateRecipeCommand) (int, error) {
	steps := make([]domain.Step, 0, len(cmd.Steps))

	for i, s := range cmd.Steps {
		step, err := domain.NewStep(s.Name, s.Instructions, i+1)
		if err != nil {
			return 0, err
		}
		steps = append(steps, step)
	}

	recipe, err := domain.NewRecipe(cmd.Name, cmd.Quantity, cmd.Unit, cmd.Difficulty, steps, cmd.Description)
	if err != nil {
		return 0, err
	}

	return uc.repo.Save(recipe)

}

func (uc *RecipeUseCase) Delete(id int) error {
	return uc.repo.Remove(id)
}
