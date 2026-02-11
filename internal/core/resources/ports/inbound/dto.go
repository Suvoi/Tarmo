package inbound

type CreateResourceCommand struct {
	Name        string
	Description string
	Price       int
	Quantity    float64
	Unit        string
}

type UpdateResourceCommand struct {
	ID          int
	Name        string
	Description string
	Price       int
	Quantity    float64
	Unit        string
}

type ResourceDTO struct {
	ID          int
	Name        string
	Description string
	Price       int
	Quantity    QuantityDTO
}

type QuantityDTO struct {
	Value float64
	Unit  UnitDTO
}

type UnitDTO struct {
	Name string
}
