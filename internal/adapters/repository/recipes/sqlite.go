package recipes

import (
	"database/sql"
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
		difficulty TEXT
	);
	`
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRecipesRepo) GetAll() ([]*domain.Recipe, error) {
	rows, err := r.db.Query(`
        SELECT id, name, description
        FROM recipes
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*domain.Recipe

	for rows.Next() {
		var rcp domain.Recipe
		if err := rows.Scan(
			&rcp.ID,
			&rcp.Name,
			&rcp.Description,
		); err != nil {
			return nil, err
		}
		recipes = append(recipes, &rcp)
	}

	return recipes, nil
}

func (r *SQLiteRecipesRepo) GetByID(id int) (*domain.Recipe, error) {
	row := r.db.QueryRow(`
			SELECT id, name, description, quantity, unit, difficulty
			FROM recipes WHERE id=?
	`, id)

	var rcp domain.Recipe
	err := row.Scan(
		&rcp.ID,
		&rcp.Name,
		&rcp.Description,
		&rcp.Quantity,
		&rcp.Unit,
		&rcp.Difficulty,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &rcp, nil
}

func (r *SQLiteRecipesRepo) Create(rcp *domain.Recipe) error {
	query := `
	        INSERT INTO recipes (name, description, quantity, unit, difficulty)
	        VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		rcp.Name,
		rcp.Description,
		rcp.Quantity,
		rcp.Unit,
		rcp.Difficulty,
	)
	if err != nil {
		return err
	}
	return nil
}
