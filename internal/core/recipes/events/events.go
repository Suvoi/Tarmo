package events

type RecipeCreated struct {
	ID   int
	Name string
}

func (e RecipeCreated) EventName() string { return "recipe.created" }

type RecipeUpdated struct {
	ID int
}

func (e RecipeUpdated) EventName() string { return "recipe.updated" }
