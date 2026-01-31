package inbound

import (
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/usecase"
)

type RecipePort interface {
	GetAll() ([]*domain.Recipe, error)
	GetByID(id int) (*domain.Recipe, error)
	Create(cmd usecase.CreateRecipeCommand) error
	Delete(id int) error
}
