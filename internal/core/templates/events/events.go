package events

type TemplateCreated struct {
	ID   int
	Name string
}

func (e TemplateCreated) EventName() string { return "template.created" }

type TemplateUpdated struct {
	ID int
}

func (e TemplateUpdated) EventName() string { return "template.updated" }
