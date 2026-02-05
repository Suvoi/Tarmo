package sqlite

import (
	"database/sql"
	"errors"
	"tarmo/internal/core/recipes"
	"tarmo/internal/core/recipes/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRecipesRepo struct {
	db *sql.DB
}

func NewSQLiteRecipesRepo(dbPath string) (*SQLiteRecipesRepo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	repo := &SQLiteRecipesRepo{db: db}
	if err := repo.init(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *SQLiteRecipesRepo) init() error {
	query := `
  CREATE TABLE IF NOT EXISTS recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    quantity INTEGER,
    unit TEXT,
    difficulty INTEGER
  );
  
  CREATE TABLE IF NOT EXISTS steps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    recipe_id INTEGER NOT NULL,
    step_order INTEGER NOT NULL,
    name TEXT NOT NULL,
    instructions TEXT,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
  );
  `
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRecipesRepo) FindAll() ([]*domain.Recipe, error) {
	// Get all recipes
	rows, err := r.db.Query(`
		SELECT id, name, description, quantity, unit, difficulty
		FROM recipes
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*domain.Recipe

	for rows.Next() {
		var (
			recipeID    int
			name        string
			description string
			quantity    int
			unit        string
			difficulty  int
		)

		if err := rows.Scan(&recipeID, &name, &description, &quantity, &unit, &difficulty); err != nil {
			return nil, err
		}

		// Get steps for this recipe
		stepRows, err := r.db.Query(`
			SELECT step_order, name, instructions
			FROM steps WHERE recipe_id = ?
			ORDER BY step_order
		`, recipeID)
		if err != nil {
			return nil, err
		}

		steps := []domain.Step{}
		for stepRows.Next() {
			var (
				order        int
				stepName     string
				instructions string
			)
			if err := stepRows.Scan(&order, &stepName, &instructions); err != nil {
				stepRows.Close()
				return nil, err
			}

			step, err := domain.NewStep(stepName, instructions, order)
			if err != nil {
				stepRows.Close()
				return nil, err
			}
			steps = append(steps, step)
		}
		stepRows.Close()

		if err := stepRows.Err(); err != nil {
			return nil, err
		}

		// Reconstruct recipe with steps
		recipe, err := domain.ReconstructRecipe(
			recipeID,
			name,
			quantity,
			unit,
			difficulty,
			steps,
			description,
		)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, rows.Err()
}

func (r *SQLiteRecipesRepo) FindByID(id int) (*domain.Recipe, error) {
	// Search recipe
	row := r.db.QueryRow(`
			SELECT id, name, description, quantity, unit, difficulty
			FROM recipes WHERE id=?
	`, id)

	// Temp variables
	var (
		recipeID    int
		name        string
		description string
		quantity    int
		unit        string
		difficulty  int
	)

	err := row.Scan(&recipeID, &name, &description, &quantity, &unit, &difficulty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, recipes.ErrRecipeNotFound // Not found
		}
		return nil, err
	}

	stepRows, err := r.db.Query(`
		SELECT step_order, name, instructions
		FROM steps WHERE recipe_id = ?
		ORDER BY step_order
	`, recipeID)
	if err != nil {
		return nil, err
	}
	defer stepRows.Close()

	steps := []domain.Step{}
	for stepRows.Next() {
		var (
			order        int
			stepName     string
			instructions string
		)
		if err := stepRows.Scan(&order, &stepName, &instructions); err != nil {
			return nil, err
		}

		step, err := domain.NewStep(stepName, instructions, order)
		if err != nil {
			return nil, err // Invalid step
		}
		steps = append(steps, step)
	}

	// Reconstruct
	return domain.ReconstructRecipe(
		recipeID,
		name,
		quantity,
		unit,
		difficulty,
		steps,
		description,
	)
}

func (r *SQLiteRecipesRepo) Save(rcp *domain.Recipe) (int, error) {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
	        INSERT INTO recipes (name, description, quantity, unit, difficulty)
	        VALUES (?, ?, ?, ?, ?)
	`
	result, err := tx.Exec(query,
		rcp.Name(),
		rcp.Description(),
		rcp.Quantity(),
		rcp.Unit(),
		rcp.Difficulty(),
	)
	if err != nil {
		return 0, err
	}

	recipeID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	stepQuery := `INSERT INTO steps (recipe_id, step_order, name, instructions) VALUES (?, ?, ?, ?)`
	for _, step := range rcp.Steps() {
		_, err := tx.Exec(stepQuery, recipeID, step.Order(), step.Name(), step.Instructions())
		if err != nil {
			return 0, err
		}
	}

	// No errors, commit
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(recipeID), nil
}

func (r *SQLiteRecipesRepo) Update(rcp *domain.Recipe) error {
	// Begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
        UPDATE recipes 
        SET name = ?, description = ?, quantity = ?, unit = ?, difficulty = ?
        WHERE id = ?
    `
	_, err = tx.Exec(query,
		rcp.Name(),
		rcp.Description(),
		rcp.Quantity(),
		rcp.Unit(),
		rcp.Difficulty(),
		rcp.ID(),
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM steps WHERE recipe_id = ?", rcp.ID())
	if err != nil {
		return err
	}

	stepQuery := `INSERT INTO steps (recipe_id, step_order, name, instructions) VALUES (?, ?, ?, ?)`
	for _, step := range rcp.Steps() {
		_, err = tx.Exec(stepQuery, rcp.ID(), step.Order(), step.Name(), step.Instructions())
		if err != nil {
			return err
		}
	}

	// No errors, commit
	return tx.Commit()
}

func (r *SQLiteRecipesRepo) Remove(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM recipes WHERE id = ?
	`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return recipes.ErrRecipeNotFound
	}

	return nil
}
