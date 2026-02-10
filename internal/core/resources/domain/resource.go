package domain

import "errors"

var (
	ErrNameRequired         = errors.New("name is required")
	ErrPriceRequired        = errors.New("price is required")
	ErrPriceGreaterThanZero = errors.New("price must be greater than 0")
)

type Resource struct {
	id          int
	name        string // [REQUIRED]
	description string
	price       int // [REQUIRED] In cents, or smallest currency unit, per 100g/100ml
}

// ===========================GETTERS===========================

func (r *Resource) ID() int             { return r.id }
func (r *Resource) Name() string        { return r.name }
func (r *Resource) Description() string { return r.description }
func (r *Resource) Price() int          { return r.price }

// ===========================CONSTRUCTORS======================

func NewResource(name string, description string, price int) (*Resource, error) {
	rsc := &Resource{
		name:        name,
		description: description,
		price:       price,
	}
	if err := rsc.Validate(); err != nil {
		return nil, err
	}
	return rsc, nil
}

func ReconstructResource(id int, name string, description string, price int) (*Resource, error) {
	rsc := &Resource{
		id:          id,
		name:        name,
		description: description,
		price:       price,
	}
	if err := rsc.Validate(); err != nil {
		return nil, err
	}
	return rsc, nil
}

// ===========================METHODS===========================

func (r *Resource) Update(name string, description string, price int) error {
	rsc, err := NewResource(name, description, price)
	if err != nil {
		return err
	}

	r.name = rsc.name
	r.description = rsc.description
	r.price = rsc.price

	return nil
}

// ===========================VALIDATORS========================

func (r *Resource) Validate() error {
	if r.name == "" {
		return ErrNameRequired
	}
	if r.price <= 0 {
		return ErrPriceGreaterThanZero
	}
	return nil
}
