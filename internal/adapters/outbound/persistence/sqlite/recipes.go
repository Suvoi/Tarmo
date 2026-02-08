package sqlite

import (
	"database/sql"
	"errors"
	"tarmo/internal/core/templates"
	"tarmo/internal/core/templates/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteTemplatesRepo struct {
	db *sql.DB
}

func NewSQLiteTemplatesRepo(dbPath string) (*SQLiteTemplatesRepo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	repo := &SQLiteTemplatesRepo{db: db}
	if err := repo.init(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *SQLiteTemplatesRepo) init() error {
	query := `
  CREATE TABLE IF NOT EXISTS templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    quantity INTEGER,
    unit TEXT,
    difficulty INTEGER
  );
  
  CREATE TABLE IF NOT EXISTS steps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    template_id INTEGER NOT NULL,
    step_order INTEGER NOT NULL,
    name TEXT NOT NULL,
    instructions TEXT,
    FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE
  );
  `
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteTemplatesRepo) FindAll() ([]*domain.Template, error) {
	// Get all templates
	rows, err := r.db.Query(`
		SELECT id, name, description, quantity, unit, difficulty
		FROM templates
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*domain.Template

	for rows.Next() {
		var (
			templateID  int
			name        string
			description string
			quantity    int
			unit        string
			difficulty  int
		)

		if err := rows.Scan(&templateID, &name, &description, &quantity, &unit, &difficulty); err != nil {
			return nil, err
		}

		// Get steps for this template
		stepRows, err := r.db.Query(`
			SELECT step_order, name, instructions
			FROM steps WHERE template_id = ?
			ORDER BY step_order
		`, templateID)
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

		// Reconstruct template with steps
		template, err := domain.ReconstructTemplate(
			templateID,
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

		templates = append(templates, template)
	}

	return templates, rows.Err()
}

func (r *SQLiteTemplatesRepo) FindByID(id int) (*domain.Template, error) {
	// Search template
	row := r.db.QueryRow(`
			SELECT id, name, description, quantity, unit, difficulty
			FROM templates WHERE id=?
	`, id)

	// Temp variables
	var (
		templateID  int
		name        string
		description string
		quantity    int
		unit        string
		difficulty  int
	)

	err := row.Scan(&templateID, &name, &description, &quantity, &unit, &difficulty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, templates.ErrTemplateNotFound // Not found
		}
		return nil, err
	}

	stepRows, err := r.db.Query(`
		SELECT step_order, name, instructions
		FROM steps WHERE template_id = ?
		ORDER BY step_order
	`, templateID)
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
	return domain.ReconstructTemplate(
		templateID,
		name,
		quantity,
		unit,
		difficulty,
		steps,
		description,
	)
}

func (r *SQLiteTemplatesRepo) Save(tmpl *domain.Template) (int, error) {
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
	        INSERT INTO templates (name, description, quantity, unit, difficulty)
	        VALUES (?, ?, ?, ?, ?)
	`
	result, err := tx.Exec(query,
		tmpl.Name(),
		tmpl.Description(),
		tmpl.Quantity(),
		tmpl.Unit(),
		tmpl.Difficulty(),
	)
	if err != nil {
		return 0, err
	}

	templateID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	stepQuery := `INSERT INTO steps (template_id, step_order, name, instructions) VALUES (?, ?, ?, ?)`
	for _, step := range tmpl.Steps() {
		_, err := tx.Exec(stepQuery, templateID, step.Order(), step.Name(), step.Instructions())
		if err != nil {
			return 0, err
		}
	}

	// No errors, commit
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(templateID), nil
}

func (r *SQLiteTemplatesRepo) Update(tmpl *domain.Template) error {
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
        UPDATE templates 
        SET name = ?, description = ?, quantity = ?, unit = ?, difficulty = ?
        WHERE id = ?
    `
	_, err = tx.Exec(query,
		tmpl.Name(),
		tmpl.Description(),
		tmpl.Quantity(),
		tmpl.Unit(),
		tmpl.Difficulty(),
		tmpl.ID(),
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM steps WHERE template_id = ?", tmpl.ID())
	if err != nil {
		return err
	}

	stepQuery := `INSERT INTO steps (template_id, step_order, name, instructions) VALUES (?, ?, ?, ?)`
	for _, step := range tmpl.Steps() {
		_, err = tx.Exec(stepQuery, tmpl.ID(), step.Order(), step.Name(), step.Instructions())
		if err != nil {
			return err
		}
	}

	// No errors, commit
	return tx.Commit()
}

func (r *SQLiteTemplatesRepo) Remove(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM templates WHERE id = ?
	`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return templates.ErrTemplateNotFound
	}

	return nil
}
