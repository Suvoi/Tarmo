package templates

import (
	"tarmo/internal/core/templates/ports/inbound"
)

func ToCreateCommand(req CreateTemplateRequestDTO) inbound.CreateTemplateCommand {
	steps := make([]inbound.StepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, inbound.StepCommand{
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}
	return inbound.CreateTemplateCommand{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		Difficulty:  req.Difficulty,
		Steps:       steps,
	}
}

func ToUpdateCommand(req UpdateTemplateRequestDTO) inbound.UpdateTemplateCommand {
	steps := make([]inbound.StepCommand, 0, len(req.Steps))
	for _, s := range req.Steps {
		steps = append(steps, inbound.StepCommand{
			Name:         s.Name,
			Instructions: s.Instructions,
		})
	}
	return inbound.UpdateTemplateCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
		Unit:        req.Unit,
		Difficulty:  req.Difficulty,
		Steps:       steps,
	}
}

func ToResponse(dto *inbound.TemplateDTO) TemplateJSONResponseDTO {
	steps := make([]StepJSONResponseDTO, len(dto.Steps))
	for i, s := range dto.Steps {
		steps[i] = StepJSONResponseDTO{
			Order:        s.Order,
			Name:         s.Name,
			Instructions: s.Instructions,
		}
	}

	return TemplateJSONResponseDTO{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Quantity:    dto.Quantity,
		Unit:        dto.Unit,
		Difficulty:  dto.Difficulty,
		Steps:       steps,
	}
}

func ToListItemResponse(r inbound.TemplateDTO) TemplateListJSONResponseDTO {
	return TemplateListJSONResponseDTO{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func ToResponseList(templates []inbound.TemplateDTO) []TemplateListJSONResponseDTO {
	res := make([]TemplateListJSONResponseDTO, 0, len(templates))
	for _, t := range templates {
		res = append(res, ToListItemResponse(t))
	}
	return res
}
