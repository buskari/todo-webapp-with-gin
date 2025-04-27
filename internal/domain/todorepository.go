package domain

type TodoRepository interface {
	FindAll() ([]*Todo, error)
	FindByID(id int) (*Todo, error)
	Create(title string) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
}
