package outbound

import (
	"tarmo/internal/core/templates/domain"
)

type TemplateRepositoryPort interface {
	FindAll() ([]*domain.Template, error)
	FindByID(id int) (*domain.Template, error)
	Save(tmpl *domain.Template) (int, error)
	Update(tmpl *domain.Template) error
	Remove(id int) error
}
