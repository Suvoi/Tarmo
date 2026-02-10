package outbound

import "tarmo/internal/core/resources/domain"

type ResourceRepositoryPort interface {
	FindAll() ([]*domain.Resource, error)
	FindByID(id int) (*domain.Resource, error)
	Save(rsc *domain.Resource) (int, error)
	Update(rsc *domain.Resource) error
	Remove(id int) error
}
