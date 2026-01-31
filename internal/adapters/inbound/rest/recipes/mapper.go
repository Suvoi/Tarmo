package recipes

import (
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/usecase"
)

func ToCommand(req CreateRecipeRequest) usecase.CreateRecipeCommand {
	steps := make([]usecase.CreateStepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, usecase.CreateStepCommand{
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}
	return usecase.CreateRecipeCommand{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		Difficulty:  req.Difficulty,
		Steps:       steps,
	}
}

func ToResponse(r *domain.Recipe) RecipeResponse {
	steps := make([]StepResponse, 0, len(r.Steps))
	for _, s := range r.Steps {
		steps = append(steps, StepResponse{
			ID:           s.ID,
			Order:        s.Order,
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}

	return RecipeResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Quantity:    r.Quantity,
		Unit:        r.Unit,
		Difficulty:  r.Difficulty,
		Steps:       steps,
	}
}

func ToListItemResponse(r *domain.Recipe) RecipeListResponse {
	return RecipeListResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func ToResponseList(recipes []*domain.Recipe) []RecipeListResponse {
	res := make([]RecipeListResponse, 0, len(recipes))
	for _, r := range recipes {
		res = append(res, ToListItemResponse(r))
	}
	return res
}
