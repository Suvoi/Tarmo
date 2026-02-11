package sqlite

import (
	"database/sql"
	"errors"
	"tarmo/internal/core/resources"
	"tarmo/internal/core/resources/domain"
	"tarmo/internal/core/shared"
)

type resourceRepository struct {
	db *sql.DB
}

func NewResourceRepository(db *SQLiteDB) *resourceRepository {
	return &resourceRepository{db: db.db}
}

func (r *resourceRepository) FindAll() ([]*domain.Resource, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price, base_quantity, base_unit
		FROM resources
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []*domain.Resource
	for rows.Next() {
		var (
			id           int
			name         string
			description  string
			price        int
			baseQuantity float64
			baseUnit     string
		)
		if err := rows.Scan(&id, &name, &description, &price, &baseQuantity, &baseUnit); err != nil {
			return nil, err
		}

		quantity, err := shared.ReconstructQuantity(baseQuantity, baseUnit)
		if err != nil {
			return nil, err
		}

		rsc, err := domain.ReconstructResource(id, name, description, price, quantity)
		if err != nil {
			return nil, err
		}
		resources = append(resources, rsc)
	}
	return resources, rows.Err()
}

func (r *resourceRepository) FindByID(id int) (*domain.Resource, error) {
	row := r.db.QueryRow(`
		SELECT id, name, description, price, base_quantity, base_unit
		FROM resources WHERE id = ?
	`, id)

	var (
		ResourceID   int
		name         string
		description  string
		price        int
		baseQuantity float64
		baseUnit     string
	)

	err := row.Scan(&ResourceID, &name, &description, &price, &baseQuantity, &baseUnit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, resources.ErrResourceNotFound // Not found
		}
		return nil, err
	}

	quantity, err := shared.ReconstructQuantity(baseQuantity, baseUnit)
	if err != nil {
		return nil, err
	}

	return domain.ReconstructResource(id, name, description, price, quantity)
}

func (r *resourceRepository) Save(resource *domain.Resource) (int, error) {
	query := `
        INSERT INTO resources (name, description, price, base_quantity, base_unit)
        VALUES (?, ?, ?, ?, ?)
    `
	result, err := r.db.Exec(query,
		resource.Name(),
		resource.Description(),
		resource.Price(),
		resource.QuantityValue(),
		resource.QuantityUnitName(),
	)
	if err != nil {
		return 0, err
	}

	resourceID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(resourceID), nil
}

func (r *resourceRepository) Update(rsc *domain.Resource) error {
	query := `
        UPDATE resources 
        SET name = ?, description = ?, price = ?, base_quantity = ?, base_unit = ?
        WHERE id = ?
    `
	_, err := r.db.Exec(query,
		rsc.Name(),
		rsc.Description(),
		rsc.Price(),
		rsc.QuantityValue(),
		rsc.QuantityUnitName(),
		rsc.ID(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *resourceRepository) Remove(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM resources WHERE id = ?
	`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return resources.ErrResourceNotFound
	}

	return nil
}
