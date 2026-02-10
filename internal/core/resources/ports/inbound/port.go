package inbound

type ResourcePort interface {
	GetAll() ([]ResourceDTO, error)
	GetByID(id int) (ResourceDTO, error)
	Create(cmd CreateResourceCommand) (int, error)
	Update(cmd UpdateResourceCommand) error
	Delete(id int) error
}
