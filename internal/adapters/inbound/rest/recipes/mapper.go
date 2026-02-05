package recipes

import (
	"tarmo/internal/core/recipes/ports/inbound"
)

func ToCreateCommand(req CreateRecipeRequestDTO) inbound.CreateRecipeCommand {
	steps := make([]inbound.StepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, inbound.StepCommand{
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

func ToUpdateCommand(req UpdateRecipeRequestDTO) inbound.UpdateRecipeCommand {
	steps := make([]inbound.StepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, inbound.StepCommand{
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}
	return inbound.UpdateRecipeCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		Difficulty:  req.Difficulty,
		Steps:       steps,
	}
}

func ToResponse(dto *inbound.RecipeDTO) RecipeJSONResponseDTO {
	steps := make([]StepJSONResponseDTO, len(dto.Steps))
	for i, s := range dto.Steps {
		steps[i] = StepJSONResponseDTO{
			Order:        s.Order,
			Name:         s.Name,
			Instructions: s.Instructions,
		}
	}

	return RecipeJSONResponseDTO{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Unit:        dto.Unit,
		Difficulty:  dto.Difficulty,
		Steps:       steps,
	}
}

func ToListItemResponse(r inbound.RecipeDTO) RecipeListJSONResponseDTO {
	return RecipeListJSONResponseDTO{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func ToResponseList(recipes []inbound.RecipeDTO) []RecipeListJSONResponseDTO {
	res := make([]RecipeListJSONResponseDTO, 0, len(recipes))
	for _, r := range recipes {
		res = append(res, ToListItemResponse(r))
	}
	return res
}
