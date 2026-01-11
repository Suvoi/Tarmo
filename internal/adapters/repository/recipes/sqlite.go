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
	return &SQLiteRecipesRepo{db: db}, nil
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
