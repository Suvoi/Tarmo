package resources

import (
	"tarmo/internal/core/resources/ports/inbound"
)

func ToCreateCommand(req CreateResourceRequestDTO) inbound.CreateResourceCommand {
	return inbound.CreateResourceCommand{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func ToUpdateCommand(req UpdateResourceRequestDTO) inbound.UpdateResourceCommand {
	return inbound.UpdateResourceCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func ToResponse(dto *inbound.ResourceDTO) ResourceJSONResponseDTO {

	return ResourceJSONResponseDTO{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func ToResponseList(resources []inbound.ResourceDTO) []ResourceJSONResponseDTO {
	res := make([]ResourceJSONResponseDTO, 0, len(resources))
	for _, r := range resources {
		res = append(res, ToResponse(&r))
	}
	return res
}
