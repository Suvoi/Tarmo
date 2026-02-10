package inbound

import "tarmo/internal/core/shared"

type CreateTemplateCommand struct {
	Name        string
	Description string
	Quantity    int
	Unit        string
	Difficulty  int
	Steps       []StepCommand
	Resources   []ResourceRefCommand
}

type UpdateTemplateCommand struct {
	ID          int
	Name        string
	Description string
	Quantity    int
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
	Quantity   int
	Unit       string
}

type TemplateDTO struct {
	ID          int
	Name        string
	Description string
	Quantity    int
	Unit        shared.Unit
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
	Quantity   int
	Unit       shared.Unit
}
