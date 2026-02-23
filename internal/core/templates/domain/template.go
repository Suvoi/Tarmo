package domain

import (
	"errors"
	"tarmo/internal/core/shared"
)

// ===========================ERRORS============================

var (
	ErrNameRequired      = errors.New("name is required")
	ErrQuantityInvalid   = errors.New("quantity must be greater than 0")
	ErrUnitRequired      = errors.New("unit must be defined")
	ErrDifficultyInvalid = errors.New("difficulty must be between 0 and 5")
	ErrNoSteps           = errors.New("template must have at least one step")
	ErrInvalidStepOrders = errors.New("step orders must be sequential starting from 1")
	ErrStepName          = errors.New("step must have a name")
	ErrStepOrder         = errors.New("step order must be greater than 0")
	ErrResourceIDInvalid = errors.New("resource ID must be greater than 0")
	ErrProductIDInvalid  = errors.New("product ID must be greater than 0")
	ErrItemTypeInvalid   = errors.New("item type must be RESOURCE or PRODUCT")
)

// ===========================MODELS============================

type Template struct {
	id          int             // [REQUIRED] [AUTO GENERATED]
	name        string          // [REQUIRED]
	description string          // [OPTIONAL]
	quantity    shared.Quantity // [REQUIRED]
	difficulty  int             // [REQUIRED]

	steps []Step // [REQUIRED] MIN 1

	inputs  []InputRequirement // [REQUIRED] MIN 1
	outputs []ProductReference // [REQUIRED] MIN 1
}

type Step struct {
	order        int    // [REQUIRED]
	name         string // [REQUIRED]
	instructions string // [OPTIONAL]
}

type InputRequirement struct {
	itemType   string          // [REQUIRED] [RESOURCE, PRODUCT]
	resourceID int             // [REQUIRED]
	quantity   shared.Quantity // [REQUIRED]
}

type ProductReference struct {
	productID int             // [REQUIRED]
	quantity  shared.Quantity // [REQUIRED]
}

// ===========================GETTERS===========================

func (t *Template) ID() int                   { return t.id }
func (t *Template) Name() string              { return t.name }
func (t *Template) Description() string       { return t.description }
func (t *Template) Quantity() shared.Quantity { return t.quantity }
func (t *Template) QuantityValue() float64    { return t.quantity.Value() }
func (t *Template) QuantityUnitName() string  { return t.quantity.Unit().Name }
func (t *Template) Difficulty() int           { return t.difficulty }

func (t *Template) Steps() []Step { return append([]Step(nil), t.steps...) }

func (t *Template) Inputs() []InputRequirement {
	return append([]InputRequirement(nil), t.inputs...)
}
func (t *Template) Outputs() []ProductReference {
	return append([]ProductReference(nil), t.outputs...)
}

func (s *Step) Order() int           { return s.order }
func (s *Step) Name() string         { return s.name }
func (s *Step) Instructions() string { return s.instructions }

func (ir *InputRequirement) ItemType() string         { return ir.itemType }
func (ir *InputRequirement) ResourceID() int          { return ir.resourceID }
func (ir *InputRequirement) QuantityValue() float64   { return ir.quantity.Value() }
func (ir *InputRequirement) QuantityUnitName() string { return ir.quantity.Unit().Name }

func (pr *ProductReference) ProductID() int           { return pr.productID }
func (pr *ProductReference) QuantityValue() float64   { return pr.quantity.Value() }
func (pr *ProductReference) QuantityUnitName() string { return pr.quantity.Unit().Name }

// ===========================CONSTRUCTORS======================

func NewInputRequirement(resourceID int, quantity float64, unitStr string) (InputRequirement, error) {
	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return InputRequirement{}, err
	}
	qty = qty.ToBase()
	rr := InputRequirement{
		resourceID: resourceID,
		quantity:   qty,
	}
	if err := rr.Validate(); err != nil {
		return InputRequirement{}, err
	}
	return rr, nil
}

func NewProductReference(itemType string, productID int, quantity float64, unitStr string) (ProductReference, error) {
	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return ProductReference{}, err
	}
	qty = qty.ToBase()
	pr := ProductReference{
		productID: productID,
		quantity:  qty,
	}
	if err := pr.Validate(); err != nil {
		return ProductReference{}, err
	}
	return pr, nil
}

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
	quantity float64,
	unitStr string,
	difficulty int,
	steps []Step,
	description string,
	resources []InputRequirement,
) (*Template, error) {

	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return nil, err
	}
	qty = qty.ToBase()
	r := &Template{
		id:          0,
		name:        name,
		quantity:    qty,
		difficulty:  difficulty,
		steps:       steps,
		description: description,
		resources:   resources,
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	return r, nil
}

func ReconstructTemplate(
	id int,
	name string,
	quantity float64,
	unitStr string,
	difficulty int,
	steps []Step,
	description string,
	resources []InputRequirement,
) (*Template, error) {

	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return nil, err
	}
	qty = qty.ToBase()
	t := &Template{
		id:          id,
		name:        name,
		quantity:    qty,
		difficulty:  difficulty,
		steps:       steps,
		description: description,
		resources:   resources,
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}

	return t, nil
}

// ===========================METHODS===========================
func (t *Template) Update(name string, quantity float64, unitStr string, difficulty int, steps []Step, description string, resources []InputRequirement) error {
	temp, err := NewTemplate(name, quantity, unitStr, difficulty, steps, description, resources)
	if err != nil {
		return err
	}

	t.name = temp.name
	t.quantity = temp.quantity
	t.difficulty = temp.difficulty
	t.steps = temp.steps
	t.description = temp.description
	t.resources = temp.resources

	return nil
}

func (t *Template) UpdateResources(resources []InputRequirement) error {
	for _, resource := range resources {
		if err := resource.Validate(); err != nil {
			return err
		}
	}
	t.resources = resources
	return nil
}

// ===========================VALIDATORS========================

func (t *Template) Validate() error {
	if t.name == "" {
		return ErrNameRequired
	}
	if t.difficulty < 0 || t.difficulty > 5 {
		return ErrDifficultyInvalid
	}
	if len(t.steps) == 0 {
		return ErrNoSteps
	}

	for i, step := range t.steps {
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

func (rr *InputRequirement) Validate() error {
	if rr.resourceID <= 0 {
		return ErrResourceIDInvalid
	}
	if rr.itemType != "RESOURCE" && rr.itemType != "PRODUCT" {
		return ErrItemTypeInvalid
	}
	return nil
}

func (pr *ProductReference) Validate() error {
	if pr.productID <= 0 {
		return ErrProductIDInvalid
	}
	return nil
}
