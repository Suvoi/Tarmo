package sqlite

import (
	"database/sql"
	"errors"
	"tarmo/internal/core/products"
	"tarmo/internal/core/products/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *SQLiteDB) *productRepository {
	return &productRepository{db: db.db}
}

func (r *productRepository) FindAll() ([]*domain.Product, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price, base_quantity, base_unit
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
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

		rsc, err := domain.ReconstructProduct(id, name, description, price, baseQuantity, baseUnit)
		if err != nil {
			return nil, err
		}
		products = append(products, rsc)
	}
	return products, rows.Err()
}

func (r *productRepository) FindByID(id int) (*domain.Product, error) {
	row := r.db.QueryRow(`
		SELECT id, name, description, price, base_quantity, base_unit
		FROM products WHERE id = ?
	`, id)

	var (
		ProductID    int
		name         string
		description  string
		price        int
		baseQuantity float64
		baseUnit     string
	)

	err := row.Scan(&ProductID, &name, &description, &price, &baseQuantity, &baseUnit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, products.ErrProductNotFound // Not found
		}
		return nil, err
	}

	return domain.ReconstructProduct(id, name, description, price, baseQuantity, baseUnit)
}

func (r *productRepository) Save(product *domain.Product) (int, error) {
	query := `
        INSERT INTO products (name, description, price, base_quantity, base_unit)
        VALUES (?, ?, ?, ?, ?)
    `
	result, err := r.db.Exec(query,
		product.Name(),
		product.Description(),
		product.Price(),
		product.QuantityValue(),
		product.QuantityUnitName(),
	)
	if err != nil {
		return 0, err
	}

	productID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(productID), nil
}

func (r *productRepository) Update(rsc *domain.Product) error {
	query := `
        UPDATE products 
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

func (r *productRepository) Remove(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM products WHERE id = ?
	`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return products.ErrProductNotFound
	}

	return nil
}
