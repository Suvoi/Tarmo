package inbound

type RecipePort interface {
	GetAll() ([]RecipeDTO, error)
	GetByID(id int) (RecipeDTO, error)
	Create(cmd CreateRecipeCommand) (int, error)
	Delete(id int) error
}
