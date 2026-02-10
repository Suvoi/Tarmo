package inbound

type CreateResourceCommand struct {
	Name        string
	Description string
	Price       int
}

type UpdateResourceCommand struct {
	ID          int
	Name        string
	Description string
	Price       int
}

type ResourceDTO struct {
	ID          int
	Name        string
	Description string
	Price       int
}
