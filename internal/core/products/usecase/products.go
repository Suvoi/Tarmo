package usecase

import (
	"tarmo/internal/core/products/domain"
	"tarmo/internal/core/products/ports/inbound"
	"tarmo/internal/core/products/ports/outbound"
	"tarmo/internal/core/shared"
)

type ProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewProductUseCase(repo outbound.ProductRepositoryPort) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) GetAll() ([]inbound.ProductDTO, error) {
	products, err := uc.repo.FindAll()
	if err != nil {
		return []inbound.ProductDTO{}, err
	}

	dtos := make([]inbound.ProductDTO, 0, len(products))
	for _, product := range products {
		if product == nil {
			continue
		}

		qtyDTO := shared.QuantityDTO{
			Value: product.QuantityValue(),
			Unit: shared.UnitDTO{
				Name: product.QuantityUnitName(),
			},
		}

		dtos = append(dtos, inbound.ProductDTO{
			ID:          product.ID(),
			Name:        product.Name(),
			Description: product.Description(),
			Price:       product.Price(),
			Quantity:    qtyDTO,
		})
	}

	return dtos, nil
}

func (uc *ProductUseCase) GetByID(id int) (inbound.ProductDTO, error) {
	product, err := uc.repo.FindByID(id)
	if err != nil {
		return inbound.ProductDTO{}, err
	}

	qtyDTO := shared.QuantityDTO{
		Value: product.QuantityValue(),
		Unit: shared.UnitDTO{
			Name: product.QuantityUnitName(),
		},
	}

	return inbound.ProductDTO{
		ID:          product.ID(),
		Name:        product.Name(),
		Description: product.Description(),
		Price:       product.Price(),
		Quantity:    qtyDTO,
	}, nil
}

func (uc *ProductUseCase) Create(cmd inbound.CreateProductCommand) (int, error) {
	product, err := domain.NewProduct(cmd.Name, cmd.Description, cmd.Price, cmd.Quantity, cmd.Unit)
	if err != nil {
		return 0, err
	}

	id, err := uc.repo.Save(product)
	return id, err
}

func (uc *ProductUseCase) Update(cmd inbound.UpdateProductCommand) error {
	product, err := uc.repo.FindByID(cmd.ID)
	if err != nil {
		return err
	}

	err = product.Update(cmd.Name, cmd.Description, cmd.Price, cmd.Quantity, cmd.Unit)
	if err != nil {
		return err
	}

	return uc.repo.Update(product)
}

func (uc *ProductUseCase) Delete(id int) error {
	return uc.repo.Remove(id)
}
