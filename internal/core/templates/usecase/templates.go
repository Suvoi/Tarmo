package usecase

import (
	"tarmo/internal/core/templates/domain"
	"tarmo/internal/core/templates/ports/inbound"
	"tarmo/internal/core/templates/ports/outbound"
)

type TemplateUseCase struct {
	repo outbound.TemplateRepositoryPort
}

func NewTemplateUseCase(repo outbound.TemplateRepositoryPort) *TemplateUseCase {
	return &TemplateUseCase{repo: repo}
}

func (uc *TemplateUseCase) GetAll() ([]inbound.TemplateDTO, error) {
	templates, err := uc.repo.FindAll()
	if err != nil {
		return []inbound.TemplateDTO{}, err
	}

	dtos := make([]inbound.TemplateDTO, 0, len(templates))
	for _, template := range templates {
		if template == nil {
			continue
		}

		dtos = append(dtos, inbound.TemplateDTO{
			ID:          template.ID(),
			Name:        template.Name(),
			Description: template.Description(),
			Quantity:    template.Quantity(),
			Unit:        template.Unit(),
			Difficulty:  template.Difficulty(),
		})
	}

	return dtos, nil
}

func (uc *TemplateUseCase) GetByID(id int) (inbound.TemplateDTO, error) {
	template, err := uc.repo.FindByID(id)
	if err != nil {
		return inbound.TemplateDTO{}, err
	}

	steps := make([]inbound.StepDTO, 0, len(template.Steps()))
	for _, step := range template.Steps() {
		steps = append(steps, inbound.StepDTO{
			Name:         step.Name(),
			Instructions: step.Instructions(),
			Order:        step.Order(),
		})
	}

	return inbound.TemplateDTO{
		ID:          template.ID(),
		Name:        template.Name(),
		Description: template.Description(),
		Quantity:    template.Quantity(),
		Unit:        template.Unit(),
		Difficulty:  template.Difficulty(),
		Steps:       steps,
	}, nil
}

func (uc *TemplateUseCase) Create(cmd inbound.CreateTemplateCommand) (int, error) {
	steps, err := mapStepsToDomain(cmd.Steps)
	if err != nil {
		return 0, err
	}

	template, err := domain.NewTemplate(cmd.Name, cmd.Quantity, cmd.Unit, cmd.Difficulty, steps, cmd.Description)
	if err != nil {
		return 0, err
	}

	id, err := uc.repo.Save(template)
	return id, err
}

func (uc *TemplateUseCase) Update(cmd inbound.UpdateTemplateCommand) error {
	template, err := uc.repo.FindByID(cmd.ID)
	if err != nil {
		return err
	}

	steps, err := mapStepsToDomain(cmd.Steps)
	if err != nil {
		return err
	}

	template.Update(cmd.Name, cmd.Quantity, cmd.Unit, cmd.Difficulty, steps, cmd.Description)

	return uc.repo.Update(template)
}

func (uc *TemplateUseCase) Delete(id int) error {
	return uc.repo.Remove(id)
}

func mapStepsToDomain(dtos []inbound.StepCommand) ([]domain.Step, error) {
	steps := make([]domain.Step, 0, len(dtos))
	for i, s := range dtos {
		step, err := domain.NewStep(s.Name, s.Instructions, i+1)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
