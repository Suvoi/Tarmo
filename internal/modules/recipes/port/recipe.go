package port

import "tarmo/internal/modules/recipes/model"

type RecipePort interface {
	GetAll() ([]*model.Recipe, error)
}
