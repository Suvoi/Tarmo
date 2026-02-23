package outbound

import "tarmo/internal/core/products/domain"

type ProductRepositoryPort interface {
	FindAll() ([]*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	Save(pct *domain.Product) (int, error)
	Update(pct *domain.Product) error
	Remove(id int) error
}
