package domain

import (
	"errors"
	"tarmo/internal/core/shared"
)

var (
	ErrNameRequired         = errors.New("name is required")
	ErrPriceGreaterThanZero = errors.New("price must be greater than 0")
	ErrUnitRequired         = errors.New("unit is required")
	ErrQuantityInvalid      = errors.New("quantity must be greater than 0")
)

// ===========================MODELS============================

type Product struct {
	id          int             // [REQUIRED] [AUTO GENERATED]
	name        string          // [REQUIRED]
	description string          // [OPTIONAL]
	price       int             // [REQUIRED] [DEFAULT 0] In cents, or smallest currency unit
	quantity    shared.Quantity // [REQUIRED] e.g. 100, 1
}

// ===========================GETTERS===========================

func (r *Product) ID() int                   { return r.id }
func (r *Product) Name() string              { return r.name }
func (r *Product) Description() string       { return r.description }
func (r *Product) Price() int                { return r.price }
func (r *Product) Quantity() shared.Quantity { return r.quantity }
func (r *Product) QuantityValue() float64    { return r.quantity.Value() }
func (r *Product) QuantityUnitName() string  { return r.quantity.Unit().Name }

// ===========================CONSTRUCTORS======================

func NewProduct(name string, description string, price int, quantity float64, unitStr string) (*Product, error) {
	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return nil, err
	}

	baseQty := qty.ToBase()

	rsc := &Product{
		name:        name,
		description: description,
		price:       price,
		quantity:    baseQty,
	}
	if err := rsc.Validate(); err != nil {
		return nil, err
	}
	return rsc, nil
}

func ReconstructProduct(id int, name string, description string, price int, quantity float64, unitStr string) (*Product, error) {
	qty, err := shared.NewQuantity(quantity, unitStr)
	if err != nil {
		return nil, err
	}

	baseQty := qty.ToBase()

	rsc := &Product{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		quantity:    baseQty,
	}
	if err := rsc.Validate(); err != nil {
		return nil, err
	}
	return rsc, nil
}

// ===========================METHODS===========================

func (r *Product) Update(name string, description string, price int, quantity float64, unitStr string) error {
	rsc, err := NewProduct(name, description, price, quantity, unitStr)
	if err != nil {
		return err
	}

	r.name = rsc.name
	r.description = rsc.description
	r.quantity = rsc.quantity

	return nil
}

// ===========================VALIDATORS========================

func (r *Product) Validate() error {
	if r.name == "" {
		return ErrNameRequired
	}
	if r.price <= 0 {
		return ErrPriceGreaterThanZero
	}
	if r.QuantityValue() <= 0 {
		return ErrQuantityInvalid
	}
	if r.QuantityUnitName() == "" {
		return ErrUnitRequired
	}
	return nil
}
