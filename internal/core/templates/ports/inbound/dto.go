package inbound

import "tarmo/internal/core/shared"

type CreateTemplateCommand struct {
	Name        string
	Description string
	Quantity    float64
	Unit        string
	Difficulty  int
	Steps       []StepCommand
	Resources   []ResourceRefCommand
}

type UpdateTemplateCommand struct {
	ID          int
	Name        string
	Description string
	Quantity    float64
	Unit        string
	Difficulty  int
	Steps       []StepCommand
	Resources   []ResourceRefCommand
}

type StepCommand struct {
	Name         string
	Instructions string
}

type ResourceRefCommand struct {
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
	Resources   []ResourceRefDTO
}

type StepDTO struct {
	Name         string
	Instructions string
	Order        int
}

type ResourceRefDTO struct {
	ResourceID int
	Quantity   shared.QuantityDTO
}
