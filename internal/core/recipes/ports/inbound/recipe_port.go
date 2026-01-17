package inbound

import "tarmo/internal/core/recipes/domain"

type RecipePort interface {
	GetAll() ([]*domain.Recipe, error)
	GetByID(id int) (*domain.Recipe, error)
	Create(rcp *domain.Recipe) error
	Delete(id int) error
}
