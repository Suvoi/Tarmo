package recipes

type StepResponse struct {
	ID           int    `json:"id"`
	Order        int    `json:"order"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}

type RecipeResponse struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	Unit        string         `json:"unit"`
	Difficulty  int            `json:"difficulty"`
	Steps       []StepResponse `json:"steps"`
}

type RecipeListResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRecipeRequest struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Quantity    int             `json:"quantity"`
	Unit        string          `json:"unit"`
	Difficulty  int             `json:"difficulty"`
	Steps       []CreateStepDTO `json:"steps"`
}

type CreateStepDTO struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}
