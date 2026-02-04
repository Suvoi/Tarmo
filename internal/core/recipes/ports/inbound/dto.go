package inbound

type CreateRecipeCommand struct {
	Name        string
	Description string
	Quantity    int
	Unit        string
	Difficulty  int
	Steps       []CreateStepCommand
}

type CreateStepCommand struct {
	Name         string
	Instructions string
}

type RecipeDTO struct {
	ID int
	Name string
	Description string
	Quantity int
	Unit string
	Difficulty int
	Steps []StepDTO
}

type StepDTO struct {
	Name string
	Instructions string
	Order int
}