package inbound

import "tarmo/internal/core/shared"

type CreateTemplateCommand struct {
	Name        string
	Description string
	Quantity    float64
	Unit        string
	Difficulty  int
	Steps       []StepCommand
	Resources   []InputRequirementCommand
}

type UpdateTemplateCommand struct {
	ID          int
	Name        string
	Description string
	Quantity    float64
	Unit        string
	Difficulty  int
	Steps       []StepCommand
	Resources   []InputRequirementCommand
}

type StepCommand struct {
	Name         string
	Instructions string
}

type InputRequirementCommand struct {
	ResourceID int
	Quantity   float64
	Unit       string
}

type TemplateDTO struct {
	ID          int
	Name        string
	Description string
	Quantity    shared.QuantityDTO
	Difficulty  int
	Steps       []StepDTO
	Resources   []InputRequirementDTO
}

type StepDTO struct {
	Name         string
	Instructions string
	Order        int
}

type InputRequirementDTO struct {
	ResourceID int
	Quantity   shared.QuantityDTO
}
