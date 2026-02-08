package inbound

type TemplatePort interface {
	GetAll() ([]TemplateDTO, error)
	GetByID(id int) (TemplateDTO, error)
	Create(cmd CreateTemplateCommand) (int, error)
	Update(cmd UpdateTemplateCommand) error
	Delete(id int) error
}
