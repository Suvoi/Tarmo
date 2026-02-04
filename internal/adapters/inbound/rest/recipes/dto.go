package recipes

type StepJSONResponse struct {
	Order        int    `json:"order"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}

type RecipeJSONResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Quantity    int                `json:"quantity"`
	Unit        string             `json:"unit"`
	Difficulty  int                `json:"difficulty"`
	Steps       []StepJSONResponse `json:"steps"`
}

type RecipeListJSONResponse struct {
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
