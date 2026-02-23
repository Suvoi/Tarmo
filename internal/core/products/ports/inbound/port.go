package inbound

type ProductPort interface {
	GetAll() ([]ProductDTO, error)
	GetByID(id int) (ProductDTO, error)
	Create(cmd CreateProductCommand) (int, error)
	Update(cmd UpdateProductCommand) error
	Delete(id int) error
}
