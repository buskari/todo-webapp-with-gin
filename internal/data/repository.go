package data

import "github.com/buskari/todo-webapp-with-gin/internal/domain"

type Repository interface {
	FindAll() ([]*domain.Todo, error)
	FindByID(id int) (*domain.Todo, error)
	Create(title string) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
}

type repository struct {
	todos []*domain.Todo
}

func New(todos []*domain.Todo) *repository {
	return &repository{
		todos: todos,
	}
}

func (r *repository) FindAll() ([]*domain.Todo, error) {
	return r.todos, nil
}

func (r *repository) FindByID(id int) (*domain.Todo, error) {
	for _, t := range r.todos {
		if t.ID() == id {
			return t, nil
		}
	}

	return nil, nil
}

func (r *repository) Create(title string) error {
	id := len(r.todos) + 1
	newTodo := domain.NewTodo(id, title)
	r.todos = append(r.todos, newTodo)
	return nil
}

func (r *repository) UpdateTitle(id int, title string) error {
	for _, t := range r.todos {
		if t.ID() == id {
			t.Title = title
		}
	}

	return nil
}

func (r *repository) Delete(id int) error {
	for _, t := range r.todos {
		if t.ID() == id && !t.IsDeleted() {
			t.ToggleDeleted()
		}
	}

	return nil
}
