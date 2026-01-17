package usecase_test

import (
	"errors"
	"tarmo/internal/core/recipes/domain"
	"tarmo/internal/core/recipes/usecase"
	"testing"
)

// mock centralizado
type mockRecipePort struct {
	allRecipes []*domain.Recipe
	single     *domain.Recipe
	errAll     error
	errSingle  error
	CreateFunc func(r *domain.Recipe) error
}

func (m *mockRecipePort) GetAll() ([]*domain.Recipe, error) {
	return m.allRecipes, m.errAll
}

func (m *mockRecipePort) GetByID(id int) (*domain.Recipe, error) {
	return m.single, m.errSingle
}

func (m *mockRecipePort) Create(r *domain.Recipe) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(r)
	}
	return nil
}

// ------------------ TESTS ------------------

func TestRecipeService_ListRecipes(t *testing.T) {
	tests := []struct {
		name      string
		mock      *mockRecipePort
		wantCount int
		wantErr   bool
	}{
		{
			name: "With recipes",
			mock: &mockRecipePort{
				allRecipes: []*domain.Recipe{
					{Name: "Pizza"},
					{Name: "Taco"},
				},
				errAll: nil,
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "No recipes",
			mock: &mockRecipePort{
				allRecipes: []*domain.Recipe{},
				errAll:     nil,
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name: "Port error",
			mock: &mockRecipePort{
				allRecipes: nil,
				errAll:     errors.New("DB error"),
			},
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := usecase.NewRecipeUseCase(tt.mock)
			got, err := service.ListRecipes()

			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != tt.wantCount {
				t.Errorf("got %d recipes, want %d", len(got), tt.wantCount)
			}
		})
	}
}

func TestRecipeService_GetRecipe(t *testing.T) {
	tests := []struct {
		name     string
		mock     *mockRecipePort
		id       int
		wantName string
		wantErr  bool
	}{
		{
			name: "Recipe found",
			mock: &mockRecipePort{
				single:    &domain.Recipe{Name: "Pizza"},
				errSingle: nil,
			},
			id:       1,
			wantName: "Pizza",
			wantErr:  false,
		},
		{
			name: "Recipe not found",
			mock: &mockRecipePort{
				single:    nil,
				errSingle: errors.New("not found"),
			},
			id:       2,
			wantName: "",
			wantErr:  true,
		},
		{
			name: "Port error",
			mock: &mockRecipePort{
				single:    nil,
				errSingle: errors.New("DB error"),
			},
			id:       3,
			wantName: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := usecase.NewRecipeUseCase(tt.mock)
			got, err := service.GetRecipe(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.Name != tt.wantName {
				t.Errorf("got %v, want %v", got.Name, tt.wantName)
			}
		})
	}
}

func TestRecipeService_CreateRecipe(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.Recipe
		wantErr bool
		wantMsg string
	}{
		{
			name:    "Valid recipe",
			input:   &domain.Recipe{Name: "Pizza", Quantity: 1, Unit: "kg", Difficulty: 3},
			wantErr: false,
		},
		{
			name:    "Missing name",
			input:   &domain.Recipe{Quantity: 1, Unit: "kg", Difficulty: 3},
			wantErr: true,
			wantMsg: "name is required",
		},
		{
			name:    "Quantity zero",
			input:   &domain.Recipe{Name: "Taco", Quantity: 0, Unit: "pcs", Difficulty: 2},
			wantErr: true,
			wantMsg: "quantity must be greater than 0",
		},
		{
			name:    "Missing unit",
			input:   &domain.Recipe{Name: "Cake", Quantity: 1, Unit: "", Difficulty: 2},
			wantErr: true,
			wantMsg: "unit must be defined",
		},
		{
			name:    "Invalid difficulty",
			input:   &domain.Recipe{Name: "Pie", Quantity: 1, Unit: "pcs", Difficulty: -1},
			wantErr: true,
			wantMsg: "difficulty must be between 0 and 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			mockPort := &mockRecipePort{
				CreateFunc: func(r *domain.Recipe) error {
					called = true
					return nil
				},
			}

			service := usecase.NewRecipeUseCase(mockPort)
			err := service.CreateRecipe(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if err.Error() != tt.wantMsg {
					t.Errorf("expected error '%s', got '%s'", tt.wantMsg, err.Error())
				}
				if called {
					t.Errorf("expected Create not to be called on invalid input")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !called {
					t.Errorf("expected Create to be called for valid input")
				}
			}
		})
	}
}
