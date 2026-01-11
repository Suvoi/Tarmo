package ports

import "tarmo/internal/core/recipes/domain"

type RecipePort interface {
	GetAll() ([]*domain.Recipe, error)
}
