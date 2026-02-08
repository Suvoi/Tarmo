package domain

import (
	"errors"
)

// Error variables
var (
	ErrNameRequired      = errors.New("name is required")
	ErrQuantityInvalid   = errors.New("quantity must be greater than 0")
	ErrUnitRequired      = errors.New("unit must be defined")
	ErrDifficultyInvalid = errors.New("difficulty must be between 0 and 5")
	ErrNoSteps           = errors.New("template must have at least one step")
	ErrInvalidStepOrders = errors.New("step orders must be sequential starting from 1")
	ErrStepName          = errors.New("step must have a name")
	ErrStepOrder         = errors.New("step order must be greater than 0")
)

type Template struct {
	id          int
	name        string // Required
	description string
	quantity    int    // Required
	unit        string // Required
	difficulty  int    // Required
	steps       []Step
}

type Step struct {
	order        int    // Required
	name         string // Required
	instructions string
}

// ===========================GETTERS===========================

func (r *Template) ID() int             { return r.id }
func (r *Template) Name() string        { return r.name }
func (r *Template) Description() string { return r.description }
func (r *Template) Quantity() int       { return r.quantity }
func (r *Template) Unit() string        { return r.unit }
func (r *Template) Difficulty() int     { return r.difficulty }
func (r *Template) Steps() []Step       { return append([]Step(nil), r.steps...) }

func (s *Step) Order() int           { return s.order }
func (s *Step) Name() string         { return s.name }
func (s *Step) Instructions() string { return s.instructions }

// ===========================CONSTRUCTORS======================

func NewStep(name, instructions string, order int) (Step, error) {
	s := Step{
		name:         name,
		instructions: instructions,
		order:        order,
	}
	if err := s.Validate(); err != nil {
		return Step{}, err
	}
	return s, nil
}

func NewTemplate(
	name string,
	quantity int,
	unit string,
	difficulty int,
	steps []Step,
	description string,
) (*Template, error) {
	r := &Template{
		id:          0,
		name:        name,
		quantity:    quantity,
		unit:        unit,
		difficulty:  difficulty,
		steps:       steps,
		description: description,
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r, nil
}

func ReconstructTemplate(
	id int,
	name string,
	quantity int,
	unit string,
	difficulty int,
	steps []Step,
	description string,
) (*Template, error) {
	r := &Template{
		id:          id,
		name:        name,
		quantity:    quantity,
		unit:        unit,
		difficulty:  difficulty,
		steps:       steps,
		description: description,
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r, nil
}

// ===========================METHODS===========================
func (r *Template) Update(name string, quantity int, unit string, difficulty int, steps []Step, description string) error {
	temp, err := NewTemplate(name, quantity, unit, difficulty, steps, description)
	if err != nil {
		return err
	}

	r.name = temp.name
	r.quantity = temp.quantity
	r.unit = temp.unit
	r.difficulty = temp.difficulty
	r.steps = temp.steps
	r.description = temp.description

	return nil
}

// ===========================VALIDATORS========================

func (r *Template) Validate() error {
	if r.name == "" {
		return ErrNameRequired
	}
	if r.quantity <= 0 {
		return ErrQuantityInvalid
	}
	if r.unit == "" {
		return ErrUnitRequired
	}
	if r.difficulty < 0 || r.difficulty > 5 {
		return ErrDifficultyInvalid
	}
	if len(r.steps) == 0 {
		return ErrNoSteps
	}

	for i, step := range r.steps {
		expectedOrder := i + 1
		if step.order != expectedOrder {
			return ErrInvalidStepOrders
		}
		if err := step.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Step) Validate() error {

	if s.order <= 0 {
		return ErrStepOrder
	}

	if s.name == "" {
		return ErrStepName
	}

	return nil
}
