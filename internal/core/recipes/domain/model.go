package domain

type Recipe struct {
	ID          int
	Name        string // Required
	Description string
	Quantity    int    // Required
	Unit        string // Required
	Difficulty  int    // Required
}
