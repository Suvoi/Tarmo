package domain

type Recipe struct {
	ID          int
	Name        string // Required
	Description string
	Quantity    int    // Required
	Unit        string // Required
	Difficulty  int    // Required
	Steps       []Step
}

type Step struct {
	ID           int
	Order        int    // Required
	Name         string // Required
	Instructions string
}
