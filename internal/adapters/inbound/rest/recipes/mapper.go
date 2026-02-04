package recipes

import (
	"tarmo/internal/core/recipes/ports/inbound"
)

func ToCommand(req CreateRecipeRequest) inbound.CreateRecipeCommand {
	steps := make([]inbound.CreateStepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, inbound.CreateStepCommand{
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}
	return inbound.CreateRecipeCommand{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		Difficulty:  req.Difficulty,
		Steps:       steps,
	}
}

func ToResponse(dto *inbound.RecipeDTO) RecipeJSONResponse {
	steps := make([]StepJSONResponse, len(dto.Steps))
	for i, s := range dto.Steps {
		steps[i] = StepJSONResponse{
			Order:        s.Order,
			Name:         s.Name,
			Instructions: s.Instructions,
		}
	}

	return RecipeJSONResponse{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Unit:        dto.Unit,
		Difficulty:  dto.Difficulty,
		Steps:       steps,
	}
}

func ToListItemResponse(r inbound.RecipeDTO) RecipeListJSONResponse {
	return RecipeListJSONResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func ToResponseList(recipes []inbound.RecipeDTO) []RecipeListJSONResponse {
	res := make([]RecipeListJSONResponse, 0, len(recipes))
	for _, r := range recipes {
		res = append(res, ToListItemResponse(r))
	}
	return res
}
