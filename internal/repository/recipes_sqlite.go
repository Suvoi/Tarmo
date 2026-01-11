package repository

import (
	"database/sql"
	"tarmo/internal/modules/recipes/model"

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

func (r *SQLiteRecipesRepo) GetAll() ([]*model.Recipe, error) {
	rows, err := r.db.Query(`
        SELECT id, name, description, quantity, unit, difficulty
        FROM recipes
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*model.Recipe

	for rows.Next() {
		var rcp model.Recipe
		if err := rows.Scan(
			&rcp.ID,
			&rcp.Name,
			&rcp.Description,
			&rcp.Quantity,
			&rcp.Unit,
			&rcp.Difficulty,
		); err != nil {
			return nil, err
		}
		recipes = append(recipes, &rcp)
	}

	return recipes, nil
}
