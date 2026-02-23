package inbound

import "tarmo/internal/core/shared"

type CreateProductCommand struct {
	Name        string
	Description string
	Price       int
	Quantity    float64
	Unit        string
}

type UpdateProductCommand struct {
	ID          int
	Name        string
	Description string
	Price       int
	Quantity    float64
	Unit        string
}

type ProductDTO struct {
	ID          int
	Name        string
	Description string
	Price       int
	Quantity    shared.QuantityDTO
}
