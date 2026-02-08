package templates

type StepJSONResponseDTO struct {
	Order        int    `json:"order"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}

type TemplateJSONResponseDTO struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Quantity    int                   `json:"quantity"`
	Unit        string                `json:"unit"`
	Difficulty  int                   `json:"difficulty"`
	Steps       []StepJSONResponseDTO `json:"steps"`
}

type TemplateListJSONResponseDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTemplateRequestDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Unit        string    `json:"unit"`
	Difficulty  int       `json:"difficulty"`
	Steps       []StepDTO `json:"steps"`
}

type UpdateTemplateRequestDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Unit        string    `json:"unit"`
	Difficulty  int       `json:"difficulty"`
	Steps       []StepDTO `json:"steps"`
}

type StepDTO struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}
