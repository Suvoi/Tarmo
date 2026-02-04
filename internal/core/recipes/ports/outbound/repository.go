package outbound

import (
	"tarmo/internal/core/recipes/domain"
)

type RecipeRepositoryPort interface {
	FindAll() ([]*domain.Recipe, error)
	FindByID(id int) (*domain.Recipe, error)
	Save(rcp *domain.Recipe) (int, error)
	Remove(id int) error
}
